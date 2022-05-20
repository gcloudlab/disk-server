package logic

import (
	"context"
	"errors"
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
	// parentId
	parentData := new(models.UserRepository)
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).
		First(parentData).Error
	if err != nil {
		return nil, err
	}
	if parentData.Id == 0 {
		return nil, errors.New("文件夹不存在")
	}

	// 更新记录的ParentId
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND deleted_at IS NULL", req.Identity).
		Update("parent_id", int64(parentData.Id)).Error
	if err != nil {
		return nil, err
	}
	return
}
