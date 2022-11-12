package logic

import (
	"context"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCommentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCommentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCommentCreateLogic {
	return &PostsCommentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCommentCreateLogic) PostsCommentCreate(req *types.PostsCommentCreateRequest, userIdentity string) (resp *types.PostsCommentCreateReply, err error) {
	new_posts_comment := &models.PostsCommentBasic{
		Identity:      helper.UUID(),
		UserIdentity:  userIdentity,
		PostsIdentity: req.PostsIdentity,
		ReplyIdentity: req.ReplyIdentity,
		ReplyName:     req.ReplyName,
		Content:       req.Content,
		Mention:       req.Mention,
		Like:          0,
		Dislike:       0,
		Read:          0,
	}

	resp = new(types.PostsCommentCreateReply)

	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "posts_identity", "reply_identity", "reply_name", "content", "mention", "like", "dislike", "created_at", "updated_at").
		Create(new_posts_comment).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Msg = "success"
	resp.Code = 0
	return
}
