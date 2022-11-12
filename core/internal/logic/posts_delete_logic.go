package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsDeleteLogic {
	return &PostsDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsDeleteLogic) PostsDelete(req *types.PostsDeleteRequest, userIdentity string) (resp *types.PostsDeleteReply, err error) {
	err = l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).
		Delete(new(models.PostsBasic)).Error

	resp = new(types.PostsDeleteReply)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
