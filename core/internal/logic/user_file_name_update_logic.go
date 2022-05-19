package logic

import (
	"context"
	"errors"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	// data := &models.UserRepository{
	// 	Name: req.Name,
	// }
	if req.Name == "" {
		err = errors.New("name is empty")
		return
	}

	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Update("name", req.Name).Error

	if err != nil {
		return
	}
	return
}
