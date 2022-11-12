package logic

import (
	"context"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateRequest, userIdentity string) (resp *types.UserUpdateReply, err error) {
	resp = new(types.UserUpdateReply)
	if req.Name != "" {
		exits := findInfoIsExits(l, "name", req.Name)
		if exits {
			resp.Msg = "用户名已存在"
			return
		}
		err = updateInfo(l, "name", req.Name, userIdentity)
		if err != nil {
			resp.Msg = "出错了"
			return
		}
		resp.Msg = "success"
	}
	if req.Email != "" {
		exits := findInfoIsExits(l, "email", req.Email)
		if exits {
			resp.Msg = "邮箱已存在"
			return
		}
		err = updateInfo(l, "email", req.Email, userIdentity)
		if err != nil {
			resp.Msg = "出错了"
			return
		}
		resp.Msg = "success"
	}
	if req.Password != "" {
		err = updateInfo(l, "password", helper.Md5(req.Password), userIdentity)
		if err != nil {
			resp.Msg = "出错了"
			return
		}
		resp.Msg = "success"
	}
	if req.Avatar != "" {
		err = updateInfo(l, "avatar", req.Avatar, userIdentity)
		if err != nil {
			resp.Msg = "出错了"
			return
		}
		resp.Msg = "success"
	}
	return
}

// 根据字段查询是否存在
func findInfoIsExits(l *UserUpdateLogic, field string, value string) (exits bool) {
	var count int64
	l.svcCtx.Engine.
		Table("user_basic").
		Where(field+" = ?", value).
		Count(&count)
	return count > 0
}

// 根据字段更新
func updateInfo(l *UserUpdateLogic, field string, value string, userIdentity string) (err error) {
	err = l.svcCtx.Engine.
		Table("user_basic").
		Where("identity = ?", userIdentity).
		Update(field, value).Error
	return
}
