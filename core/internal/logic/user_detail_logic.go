package logic

import (
	"context"

	"gcloud/core/helper"
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

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest, authorization string) (resp *types.UserDetailReply, err error) {
	resp = &types.UserDetailReply{}
	userClaim, err := helper.AnalyzeToken(authorization)
	if err != nil {
		resp.Msg = "expired token"
		return
	}
	user_detail := new(models.UserBasic)

	l.svcCtx.Engine.
		Where("name = ?", userClaim.Name).
		First(user_detail)
	if user_detail.Id == 0 {
		resp.Msg = "not found"
		return
	}

	resp.Name = user_detail.Name
	resp.Email = user_detail.Email
	resp.Identity = user_detail.Identity
	resp.Avatar = user_detail.Avatar
	resp.Capacity = user_detail.Capacity
	resp.CreatedAt = user_detail.CreatedAt.String()
	resp.Msg = "success"
	return
}
