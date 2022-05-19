package logic

import (
	"context"
	"time"

	"gcloud/core/define"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	usrFile := make([]*types.UserFile, 0)
	var cnt int64
	resp = new(types.UserFileListReply)

	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	err = l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Limit(size).
		Offset(offset).
		Find(&usrFile).Error

	if err != nil {
		return
	}

	// 查询总数
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).
		Count(&cnt).Error
	if err != nil {
		return
	}

	resp.List = usrFile
	resp.Count = cnt
	return
}
