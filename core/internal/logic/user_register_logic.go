package logic

import (
	"context"
	"gcloud/core/helper"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	resp = new(types.UserRegisterReply)
	// 判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		resp.Msg = "无效验证码"
		return
	}
	if code != req.Code {
		resp.Msg = "验证码错误"
		return
	}

	// 判断用户名是否已存在
	var count int64
	err = l.svcCtx.Engine.
		Table("user_basic").
		Where("name = ?", req.Name).
		Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		resp.Msg = "用户名已存在"
		return
	}

	// 入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: helper.Md5(req.Password),
	}
	// fix: 需指定添加字段 (Select())，不推荐使用 Omit()
	err = l.svcCtx.Engine.
		Select("identity", "name", "email", "password", "created_at", "updated_at").
		Create(user).Error
	if err != nil {
		return
	}
	resp.Msg = "注册成功"
	return
}
