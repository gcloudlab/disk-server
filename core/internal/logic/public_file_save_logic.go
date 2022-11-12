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

type PublicFileSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileSaveLogic {
	return &PublicFileSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileSaveLogic) PublicFileSave(req *types.PublicRepositorySaveRequest, UserIdentity string) (resp *types.PublicRepositorySaveReply, err error) {
	usr := &models.PublicRepository{
		Identity:           helper.UUID(),
		UserIdentity:       UserIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               req.Name,
		Ext:                req.Ext,
	}

	resp = new(types.PublicRepositorySaveReply)
	var Size struct {
		TotalSize int `json:"total_size"`
	}
	l.svcCtx.Engine.
		Table("public_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("public_repository.user_identity = ? AND public_repository.deleted_at IS NULL", UserIdentity).
		Joins("left join repository_pool on public_repository.repository_identity = repository_pool.identity").
		Take(&Size)
	if UserIdentity != "USER_1" && Size.TotalSize >= define.PublicRepositoryMaxSize {
		resp.Msg = "容量不足"
		return
	}

	var count int64
	err = l.svcCtx.Engine.
		Table("public_repository").
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
