package user_service

import (
	"ARPostcard_server/biz/consts"
	"ARPostcard_server/biz/dao"
	"ARPostcard_server/biz/model"
	"ARPostcard_server/biz/utils"
	"context"
	"github.com/RanFeng/ierror"
)

type UserLoginByUsernameReq struct {
	Username string `json:"user_name" query:"user_name" binding:"required"`
	Password string `json:"password" query:"password" binding:"required"`
}

type UserRegisterReq struct {
	Username string `json:"user_name" query:"user_name" binding:"required"`
	Password string `json:"password" query:"password" binding:"required"`
}

type UserListReq struct {
	PageNum  int `json:"pageNum" query:"pageNum"`
	PageSize int `json:"pageSize" query:"pageSize"`
}

type UserCreateReq struct {
	Username string `json:"user_name" query:"user_name" binding:"required"`
	Password string `json:"password" query:"password" binding:"required"`
	NickName string `json:"nick_name" query:"nick_name"`
	Role     int    `json:"role" query:"role"`
}

type UserUpdateReq struct {
	ID       int    `json:"id" query:"id" binding:"required"`
	NickName string `json:"nick_name" query:"nick_name"`
	Password string `json:"password" query:"password"`
	Role     int    `json:"role_id" query:"role_id"`
}

// UserDeleteReq 按照一个完整的列表删除
type UserDeleteReq struct {
	IDList []int `json:"id_list" query:"id_list" binding:"required"`
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
		UserID: uint(user.ID),
		Role:   int(user.Role),
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

func UserList(ctx context.Context, pageNum int, pageSize int) ([]*model.UserReq, int64, error) {
	users, total, err := dao.PGetUser(ctx, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func UserCreate(ctx context.Context, req UserCreateReq) error {
	// 查询用户是否存在
	user, err := dao.QueryUserByUserName(ctx, req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return ierror.NewIError(consts.UserExist, "用户已存在")
	}

	// 密码加密
	pwd := utils.HashPwd(req.Password)
	if err != nil {
		return err
	}

	// 创建用户
	err = dao.CreateUser(ctx, req.Username, pwd)
	if err != nil {
		return err
	}

	return nil
}

func UserUpdate(ctx context.Context, req UserUpdateReq) error {
	// 查询用户是否存在
	user, err := dao.QueryUserByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if user == nil {
		return ierror.NewIError(consts.UserNotExist, "用户不存在")
	}

	// 密码加密
	pwd := utils.HashPwd(req.Password)
	if err != nil {
		return err
	}

	user.Password = pwd
	user.Role = consts.Role(req.Role)
	// 更新用户，用户名不允许修改
	err = dao.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func UserDelete(ctx context.Context, idList []int) error {
	err := dao.DeleteUsers(ctx, idList)
	if err != nil {
		return err
	}
	return nil
}
