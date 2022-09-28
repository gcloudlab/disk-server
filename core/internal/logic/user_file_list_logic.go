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
	deletedFile := make([]*types.DeletedUserFile, 0)
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
	// offset := (page - 1) * size

	// TODO 按文件名查询
	// id := req.Id =》 parent_id
	// if id == 0 {
	// 	id = -1
	// }

	err = l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.updated_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("user_identity = ? ", userIdentity).
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Find(&usrFile).Error
	// Limit(size).
	// Offset(offset).
	if err != nil {
		resp.Msg = "error"
		return
	}

	err = l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.deleted_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("user_identity = ? ", userIdentity).
		Where("user_repository.deleted_at IS NOT NULL").
		// Order("user_repository.deleted_at desc").
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Find(&deletedFile).Error

	if err != nil {
		resp.Msg = "error"
		return
	}

	// 查询总数
	err = l.svcCtx.Engine.
		Table("user_repository").
		// TODO parent_id = ? AND
		Where("user_identity = ? AND deleted_at IS NULL", userIdentity).
		Count(&cnt).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = usrFile
	resp.DeletedList = deletedFile
	resp.Count = cnt
	resp.Msg = "success"
	return
}
