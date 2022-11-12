package logic

import (
	"context"
	"time"

	"gcloud/core/define"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileListLogic {
	return &PublicFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileListLogic) PublicFileList(req *types.PublicFileListRequest) (resp *types.PublicFileListReply, err error) {
	publicFile := make([]*types.PublicFile, 0)
	var cnt int64
	resp = new(types.PublicFileListReply)

	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	// offset := (page - 1) * size

	err = l.svcCtx.Engine.
		Table("public_repository").
		Select("public_repository.id, public_repository.parent_id, public_repository.identity, "+
			"public_repository.repository_identity, public_repository.ext, public_repository.updated_at,"+
			"public_repository.name, repository_pool.path, repository_pool.size, user_basic.name as owner").
		Where("public_repository.deleted_at = ? OR public_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Joins("left join repository_pool on public_repository.repository_identity = repository_pool.identity").
		Joins("left join user_basic on public_repository.user_identity = user_basic.identity").
		Find(&publicFile).Error
	// Limit(size).
	// Offset(offset).

	if err != nil {
		resp.Msg = "error"
		return
	}

	// 查询总数
	err = l.svcCtx.Engine.
		Table("public_repository").
		// TODO parent_id = ? AND
		Where("deleted_at IS NULL").
		Count(&cnt).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = publicFile
	resp.Count = cnt
	resp.Msg = "success"
	return
}
