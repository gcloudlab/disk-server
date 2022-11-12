package logic

import (
	"context"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFolderCreateLogic {
	return &PublicFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFolderCreateLogic) PublicFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	if req.Name == "" {
		resp.Msg = "name is empty"
		return
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err = l.svcCtx.Engine.
		Table("public_repository").
		Where("name = ? AND parent_id = ? AND user_identity = ?", req.Name, req.ParentId, userIdentity).
		Count(&cnt).Error

	resp = new(types.UserFolderCreateReply)
	if err != nil {
		resp.Msg = "error"
		return
	}
	if cnt > 0 {
		resp.Msg = "exits"
		return
	}

	// 创建文件夹
	data := &models.PublicRepository{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	err = l.svcCtx.Engine.
		Table("public_repository").
		Select("identity", "name", "user_identity", "parent_id", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
