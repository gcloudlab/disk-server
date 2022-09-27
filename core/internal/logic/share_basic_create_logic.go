package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	resp = new(types.ShareBasicCreateReply)
	idna := helper.UUID()
	usr := new(models.UserRepository)
	err = l.svcCtx.Engine.
		Where("identity = ?", req.UserRepositoryIdentity).
		First(usr).Error
	if err != nil {
		resp.Msg = "error"
		return
	}
	if usr.Id == 0 {
		resp.Msg = "user resource not found"
		return
	}

	data := &models.ShareBasic{
		Identity:               idna,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     usr.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
		Desc:                   req.Desc,
	}
	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "repository_identity", "user_repository_identity", "expired_time", "desc", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Identity = idna
	resp.Msg = "success"
	return
}
