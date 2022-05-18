package logic

import (
	"context"
	"errors"
	"gcloud/core/models"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 从数据库中查询当前用户
	user := new(models.UserBasic)
	has, err := models.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, 10000)
	if err != nil {
		return nil, err
	}

	resp = new(types.LoginReply)
	resp.Token = token
	return
}
