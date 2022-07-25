package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	rp := &models.RepositoryPool{
		Identity: helper.UUID(),
		Name:     req.Name,
		Hash:     req.Hash,
		Path:     req.Path,
		Ext:      req.Ext,
		Size:     req.Size,
	}
	resp = new(types.FileUploadReply)
	err = l.svcCtx.Engine.
		Select("identity", "name", "hash", "path", "ext", "size", "created_at", "updated_at").
		Create(rp).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	resp.Msg = "success"
	return
}
