package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareListLogic {
	return &UserShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserShareListLogic) UserShareList(req *types.UserShareListRequest, UserIdentity string) (resp *types.UserShareListReply, err error) {
	shareFile := make([]*types.ShareBasicDetailReply, 0)
	resp = new(types.UserShareListReply)

	err = l.svcCtx.Engine.
		Table("share_basic").
		Select("share_basic.identity, share_basic.repository_identity, user_repository.name, repository_pool.ext, "+
			"repository_pool.path, repository_pool.size, share_basic.click_num, share_basic.desc, "+
			"user_basic.name as owner, user_basic.avatar, share_basic.expired_time, share_basic.updated_at").
		Joins("LEFT JOIN repository_pool ON repository_pool.identity = share_basic.repository_identity").
		Joins("LEFT JOIN user_repository ON user_repository.identity = share_basic.user_repository_identity").
		Joins("left join user_basic on share_basic.user_identity = user_basic.identity").
		Where("share_basic.user_identity = ? ", UserIdentity).
		Where("share_basic.deleted_at IS NULL").
		Order("share_basic.updated_at desc").
		Find(&shareFile).Error

	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = shareFile
	resp.Msg = "success"
	resp.Code = 0

	return
}
