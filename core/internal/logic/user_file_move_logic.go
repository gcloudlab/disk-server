package logic

import (
	"context"
	"gcloud/core/models"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	resp = new(types.UserFileMoveReply)
	// parentId
	parentData := new(models.UserRepository)
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).
		First(parentData).Error
	if err != nil {
		resp.Msg = "error"
		return
	}
	if parentData.Id == 0 {
		resp.Msg = "文件夹不存在"
		return
	}

	// 更新记录的ParentId
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND deleted_at IS NULL", req.Identity).
		Update("parent_id", int64(parentData.Id)).Error
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
