package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	resp = new(types.UserFileNameUpdateReply)
	if req.Name == "" {
		resp.Msg = "文件名为空"
		return
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ?", req.Name).
		Where("parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Identity).
		Where("deleted_at IS NULL").
		Count(&cnt).Error

	if err != nil {
		resp.Msg = "error"
		return
	}
	if cnt > 0 {
		resp.Msg = "文件名已存在"
		return
	}

	// 更新文件名
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Update("name", req.Name).Error

	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
