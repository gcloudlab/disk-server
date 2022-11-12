package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCommentsLogic {
	return &PostsCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCommentsLogic) PostsComments(req *types.PostsCommentsRequest) (resp *types.PostsCommentsReply, err error) {
	resp = new(types.PostsCommentsReply)
	postsCommentList := make([]*types.PostsCommentsItem, 0)

	err = l.svcCtx.Engine.
		Table("posts_comment_basic").
		Select("posts_comment_basic.identity, user_basic.name as owner, posts_comment_basic.user_identity as owner_identity, user_basic.avatar, "+
			"posts_comment_basic.content, posts_comment_basic.mention, posts_basic.title as posts_name, "+
			"posts_comment_basic.like, posts_comment_basic.dislike, posts_comment_basic.updated_at, "+
			"posts_comment_basic.reply_identity, posts_comment_basic.reply_name, posts_comment_basic.posts_identity").
		Joins("left join user_basic on posts_comment_basic.user_identity = user_basic.identity").
		Joins("left join posts_basic on posts_basic.identity = posts_comment_basic.posts_identity").
		Where("posts_comment_basic.posts_identity = ?", req.PostsIdentity).
		Where("posts_comment_basic.deleted_at IS NULL").
		Find(&postsCommentList).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = postsCommentList
	resp.Msg = "success"
	resp.Code = 0
	return
}
