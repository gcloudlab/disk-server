package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, UserIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	// 用户新增文件
	usr := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       UserIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               req.Name,
		Ext:                req.Ext,
	}

	resp = new(types.UserRepositorySaveReply)
	var Size struct {
		TotalSize int `json:"total_size"`
	}
	l.svcCtx.Engine.
		Table("user_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("user_repository.user_identity = ? AND user_repository.deleted_at IS NULL", UserIdentity).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Take(&Size)
	if Size.TotalSize >= define.UserRepositoryMaxSize {
		resp.Msg = "容量不足"
		return
	}

	var count int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", req.Name, req.ParentId, UserIdentity).
		Count(&count).Error
	if count > 0 {
		resp.Msg = "exist"
		resp.Code = 405
		return
	}

	err = l.svcCtx.Engine.
		Select("identity", "parent_id", "user_identity", "repository_identity", "name", "ext", "created_at", "updated_at").
		Create(usr).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Msg = "success"
	resp.Code = 200
	return
}
