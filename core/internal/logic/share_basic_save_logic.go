package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	// logic：其他用户保存分享文件
	resp = new(types.ShareBasicSaveReply)
	// 获取资源详情 from repository_pool
	rp := new(models.RepositoryPool)
	err = l.svcCtx.Engine.
		Table("repository_pool").
		Where("identity = ?", req.RepositoryIdentity).
		First(rp).Error
	if err != nil {
		resp.Msg = "error"
		return
	}
	if rp.Id == 0 {
		resp.Msg = "资源不存在"
		return
	}

	var Size struct {
		TotalSize int `json:"total_size"`
	}
	l.svcCtx.Engine.
		Table("user_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("user_repository.user_identity = ? AND user_repository.deleted_at IS NULL", userIdentity).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Take(&Size)
	if Size.TotalSize >= define.UserRepositoryMaxSize {
		resp.Msg = "容量不足"
		return
	}

	var count int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", rp.Name, req.ParentId, userIdentity).
		Count(&count).Error
	if count > 0 {
		resp.Msg = "exist"
		resp.Code = 405
		return
	}

	// 资源保存 to user_repository
	usr := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}

	err = l.svcCtx.Engine.
		Select("identity", "parent_id", "user_identity", "repository_identity", "name", "ext", "created_at", "updated_at").
		Create(usr).Error
	if err != nil {
		resp.Msg = "save error"
		return
	}

	resp.Identity = usr.Identity
	resp.Msg = "success"
	return
}
