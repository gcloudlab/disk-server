package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GongDeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGongDeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GongDeUpdateLogic {
	return &GongDeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GongDeUpdateLogic) GongDeUpdate(req *types.GongDeUpdateRequest) (resp *types.GongDeUpdateReply, err error) {
	resp = new(types.GongDeUpdateReply)

	err = l.svcCtx.Engine.
		Table("gongde_basic").
		Exec("UPDATE gongde_basic SET count = count + ? where id = 1", req.CurrentCount).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	err = l.svcCtx.Engine.
		Table("gongde_basic").
		Select("count").
		First(resp).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Msg = "success"
	return
}
