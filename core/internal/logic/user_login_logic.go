package logic

import (
	"context"
	"errors"
	"gcloud/core/define"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 从数据库中查询当前用户
	user := new(models.UserBasic)
	l.svcCtx.Engine.
		Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).
		First(&user)
	// has, err := models.Engine.Where("name = ? AND password = ?", req.Name, req.Password).Get(user)

	if user.Id == 0 {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成普通 token1
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	// 生成用于刷新 token1 的 token2
	// 当 token1 失效后，使用 token2 生成新 token1
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}

	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
