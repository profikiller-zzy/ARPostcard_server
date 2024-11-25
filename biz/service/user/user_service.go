package user_service

import (
	"context"
	"github.com/RanFeng/ierror"
	"uav/biz/consts"
	"uav/biz/dao"
	"uav/biz/model"
	"uav/biz/utils"
)

type UserLoginByUsernameReq struct {
	Username string `json:"username" query:"username" binding:"required"`
	Password string `json:"password" query:"password" binding:"required"`
}

type UserRegisterReq struct {
	Username string `json:"username" query:"username" binding:"required"`
	Password string `json:"password" query:"password" binding:"required"`
}

// UserLogin 用户登录
func UserLogin(ctx context.Context, username, password string) (string, error) {
	var user *model.User
	var err error
	user, err = dao.QueryUserByUserName(ctx, username)
	if err != nil {
		return "", err
	}

	// 验证密码
	if !utils.CheckPwd(user.Password, password) {
		return "", ierror.NewIError(consts.PasswordWrong, "密码错误")
	}

	payLoad := utils.JwtPayLoad{
		UserID:   uint(user.ID),
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Role:     int(user.Role),
	}

	// 生成token
	token, err := utils.GenToken(payLoad)

	return token, err
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, username, password string) error {
	// 查询用户是否存在
	user, err := dao.QueryUserByUserName(ctx, username)
	if err != nil {
		return err
	}
	if user != nil {
		return ierror.NewIError(consts.UserExist, "用户已存在")
	}

	// 密码加密
	pwd := utils.HashPwd(password)
	if err != nil {
		return err
	}

	// 创建用户
	err = dao.CreateUser(ctx, username, pwd)
	if err != nil {
		return err
	}

	return nil
}
