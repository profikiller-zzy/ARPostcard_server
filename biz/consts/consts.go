package consts

import "github.com/RanFeng/ierror"

const (
	DBError = -200001 - iota
	RedisError
)

const (
	Service = iota + 10000
	Params
	AuthorizationFailed
	UserAlreadyExist
	UserIsNotExist
	TokenInvalid
	UserNameExist
	PasswordWrong
	PhoneNumInvalid
	NoLogin
	UserExist
	UserNotExist
)

var (
	TokenIsInvalid = ierror.NewIError(TokenInvalid, "Token is invalid")
	ErrNoLogin     = ierror.NewIError(NoLogin, "未登录")
	ErrUserExist   = ierror.NewIError(UserExist, "用户已存在")
)
