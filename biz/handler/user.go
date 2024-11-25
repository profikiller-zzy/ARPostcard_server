package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	user_service "uav/biz/service/user"
	"uav/biz/utils"
)

// UserLogin 用户登录
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_service.UserLoginByUsernameReq

	err = c.BindAndValidate(&req) // 传递json
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	token, err := user_service.UserLogin(ctx, req.Username, req.Password)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	utils.RespOK(ctx, c, map[string]string{"token": token})
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_service.UserRegisterReq

	err = c.BindAndValidate(&req) // 传递json
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	err = user_service.UserRegister(ctx, req.Username, req.Password)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	utils.RespOK(ctx, c, nil)
}
