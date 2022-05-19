package logic

import (
	"context"
	"errors"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	resp = &types.UserDetailReply{}
	user_detail := new(models.UserBasic)

	models.Engine.Where("identity = ?", req.Identity).First(user_detail)
	if user_detail.Id == 0 {
		return nil, errors.New("user not fonud")
	}

	resp.Name = user_detail.Name
	resp.Email = user_detail.Email
	return
}
