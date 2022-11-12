package logic

import (
	"context"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareStatisticsLogic {
	return &ShareStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareStatisticsLogic) ShareStatistics(req *types.ShareStatisticsRequest) (resp *types.ShareStatisticsReply, err error) {
	resp = &types.ShareStatisticsReply{}

	var share_count int64
	err = l.svcCtx.Engine.
		Table("share_basic").
		Count(&share_count).Error
	if err != nil {
		resp.Msg = "error"
		return
	}

	var Num struct {
		ClickNum int `json:"click_num"`
	}
	err = l.svcCtx.Engine.
		Table("share_basic").
		Select("sum(share_basic.click_num) as click_num").
		Take(&Num).Error
	// Where("share_basic.deleted_at IS NULL").
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.ShareCount = int(share_count)
	resp.ClickNum = Num.ClickNum
	resp.Msg = "success"
	resp.Code = 0
	return
}
