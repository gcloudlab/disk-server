package logic

import (
	"context"
	"errors"
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
	idna := helper.UUID()
	usr := new(models.UserRepository)
	err = l.svcCtx.Engine.
		Where("identity = ?", req.UserRepositoryIdentity).
		First(usr).Error
	if err != nil {
		return
	}
	if usr.Id == 0 {
		return nil, errors.New("user repository not found")
	}

	data := &models.ShareBasic{
		Identity:               idna,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     usr.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "repository_identity", "user_repository_identity", "expired_time", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		return
	}

	resp = &types.ShareBasicCreateReply{
		Identity: idna,
	}
	return
}
