package logic

import (
	"context"
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
		return
	}

	resp.Identity = usr.Identity
	resp.Msg = "success"
	return
}
