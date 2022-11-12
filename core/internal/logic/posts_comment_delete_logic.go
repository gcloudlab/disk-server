package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCommentDeleteLogic {
	return &PostsCommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCommentDeleteLogic) PostsCommentDelete(req *types.PostsCommentDeleteRequest, userIdentity string) (resp *types.PostsCommentDeleteReply, err error) {
	err = l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).
		Delete(new(models.PostsCommentBasic)).Error

	resp = new(types.PostsCommentDeleteReply)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
