package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsUpdateLogic {
	return &PostsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsUpdateLogic) PostsUpdate(req *types.PostsUpdateRequest, userIdentity string) (resp *types.PostsUpdateReply, err error) {
	resp = new(types.PostsUpdateReply)
	if req.Title == "" {
		resp.Msg = "empty title is not surppot"
		return
	}

	var cnt_title int64
	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("title = ? AND identity != ? AND user_identity = ? AND deleted_at IS NULL", req.Title, req.Identity, userIdentity).
		Count(&cnt_title).Error
	if err != nil {
		resp.Msg = "error"
		return
	}
	if cnt_title > 0 {
		resp.Msg = "exits"
		resp.Code = 405
		return
	}

	// 更新
	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Updates(map[string]interface{}{"title": req.Title, "content": req.Content, "tags": req.Tags, "mention": req.Mention, "cover": req.Cover}).Error

	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	resp.Code = 0
	return
}
