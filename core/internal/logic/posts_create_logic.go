package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCreateLogic {
	return &PostsCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCreateLogic) PostsCreate(req *types.PostsCreateRequest, userIdentity string) (resp *types.PostsCreateReply, err error) {
	new_posts := &models.PostsBasic{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		Title:        req.Title,
		Tags:         req.Tags,
		Content:      req.Content,
		Mention:      req.Mention,
		Cover:        req.Cover,
		ClickNum:     0,
	}

	resp = new(types.PostsCreateReply)

	var count int64
	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("title = ? AND user_identity = ? AND deleted_at IS NULL", req.Title, userIdentity).
		Count(&count).Error
	if count > 0 {
		resp.Msg = "exist"
		resp.Code = 405
		return
	}

	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "title", "tags", "content", "mention", "cover", "click_num", "created_at", "updated_at").
		Create(new_posts).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Msg = "success"
	resp.Code = 0
	return
}
