package logic

import (
	"context"
	"gcloud/core/define"
	"time"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	resp = new(types.MailCodeSendReply)
	// 1 邮箱未注册
	var count int64
	// 1.1 查询当前邮箱是否在数据库中
	err = l.svcCtx.Engine.
		Table("user_basic").
		Where("email = ?", req.Email).
		Count(&count).Error
	if err != nil {
		return
	}

	if count > 0 {
		resp.Msg = "registered"
		return
	}

	// 1.2 生成验证码
	code := helper.RandCode()
	// 1.3 存储验证码 -> redis
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	// 1.4 发送邮件验证码
	err = helper.SendMailCode(req.Email, code)
	if err != nil {
		return nil, err
	}

	return
}
