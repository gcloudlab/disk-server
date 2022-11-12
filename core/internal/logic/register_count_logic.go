package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCountLogic {
	return &RegisterCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterCountLogic) RegisterCount(req *types.RegisterCountRequest) (resp *types.RegisterCountReply, err error) {
	resp = &types.RegisterCountReply{}
	var count int64
	err = l.svcCtx.Engine.
		Table("user_basic").
		Count(&count).Error
	if err != nil {
		resp.Msg = "出错了"
		return
	}

	resp.Msg = "success"
	resp.Count = count
	return
}
