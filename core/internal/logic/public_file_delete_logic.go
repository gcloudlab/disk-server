package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileDeleteLogic {
	return &PublicFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileDeleteLogic) PublicFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteReply, err error) {
	err = l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).
		Delete(new(models.PublicRepository)).Error

	resp = new(types.UserFileDeleteReply)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
