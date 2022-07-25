package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	rp := new(models.RepositoryPool)
	resp = new(types.FileUploadPrepareReply)

	l.svcCtx.Engine.
		Where("hash = ?", req.Md5).
		First(rp)

	if rp.Id != 0 {
		// 文件已存在，秒传成功
		resp.Identity = rp.Identity
	} else {
		// 文件不存在，获取文件的 UploadID、key，执行分片上传
		key, uploadId, err := helper.CosInitPart(req.Ext)
		if err != nil {
			resp.Msg = "error"
			return resp, err
		}
		resp.Key = key
		resp.UploadId = uploadId
		resp.Msg = "success"
	}

	return
}
