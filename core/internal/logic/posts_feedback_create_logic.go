package logic

import (
	"context"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsFeedbackCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsFeedbackCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsFeedbackCreateLogic {
	return &PostsFeedbackCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsFeedbackCreateLogic) PostsFeedbackCreate(req *types.PostsFeedbackCreateRequest, userIdentity string) (resp *types.PostsFeedbackCreateReply, err error) {
	new_fb := &models.PostsFeedback{
		Identity:      helper.UUID(),
		UserIdentity:  userIdentity,
		PostsIdentity: req.PostsIdentity,
		Type:          req.Type,
		Count:         1,
		Read:          0,
	}

	resp = new(types.PostsFeedbackCreateReply)

	var count int64
	err = l.svcCtx.Engine.
		Table("posts_fb").
		Where("type = ? AND count = 1 AND posts_identity = ? AND user_identity = ? AND deleted_at IS NULL", req.Type, req.PostsIdentity, userIdentity).
		Count(&count).Error
	if count > 0 {
		l.svcCtx.Engine.
			Table("posts_fb").
			Exec("UPDATE posts_fb SET count = 0 where type = ? AND posts_identity = ? AND user_identity = ?", req.Type, req.PostsIdentity, userIdentity)
		resp.Msg = "exist"
		resp.Code = 405
	} else {
		err = l.svcCtx.Engine.
			Select("identity", "user_identity", "posts_identity", "type", "count", "created_at", "updated_at").
			Create(new_fb).Error
		if err != nil {
			resp.Msg = "error"
			return
		}
	}

	var ilike int64
	err = l.svcCtx.Engine.
		Table("posts_fb").
		Where("posts_identity = ? AND type = 'ilike' AND count = 1 AND deleted_at IS NULL", req.PostsIdentity).
		Count(&ilike).
		Error
	var dislike int64
	err = l.svcCtx.Engine.
		Table("posts_fb").
		Where("posts_identity = ? AND type = 'dislike' AND count = 1 AND deleted_at IS NULL", req.PostsIdentity).
		Count(&dislike).
		Error
	var collect int64
	err = l.svcCtx.Engine.
		Table("posts_fb").
		Where("posts_identity = ? AND type = 'collect' AND count = 1 AND deleted_at IS NULL", req.PostsIdentity).
		Count(&collect).
		Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Ilike = int(ilike)
	resp.Dislike = int(dislike)
	resp.Collect = int(collect)
	resp.Msg = "success"
	return
}
