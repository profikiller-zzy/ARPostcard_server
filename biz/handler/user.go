package handler

import (
	user_service "ARPostcard_server/biz/service/user"
	"ARPostcard_server/biz/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
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

// UserList 用户列表
func UserList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_service.UserListReq

	err = c.BindAndValidate(&req) // 传递json
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	users, total, err := user_service.UserList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	data := map[string]interface{}{
		"list":  users,
		"total": total,
	}
	utils.RespOK(ctx, c, data)
}

// UserCreate 创建用户
func UserCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_service.UserCreateReq

	err = c.BindAndValidate(&req) // 传递json
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	err = user_service.UserCreate(ctx, req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	utils.RespOK(ctx, c, nil)
}

// UserUpdate 更新用户
func UserUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_service.UserUpdateReq

	err = c.BindAndValidate(&req) // 传递json
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	err = user_service.UserUpdate(ctx, req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	utils.RespOK(ctx, c, nil)
}

// UserDelete 删除用户
func UserDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user_service.UserDeleteReq

	err = c.BindAndValidate(&req) // 传递json
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	err = user_service.UserDelete(ctx, req.IDList)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	utils.RespOK(ctx, c, nil)
}
