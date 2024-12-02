package model

import "ARPostcard_server/biz/consts"

// User 用户表
type User struct {
	MODEL
	UserName string      `gorm:"size:36" json:"user_name"`                  // 用户名
	Password string      `gorm:"size:128" json:"password"`                  // 密码
	Email    string      `gorm:"size:128" json:"email,select(info)"`        // 邮箱
	Role     consts.Role `gorm:"size:4;default:1" json:"role,select(info)"` // 权限  1 管理员  2 普通用户  3 游客
}

func (User) TableName() string {
	return "users"
}

// UserReq 用于给到前端的返回值
type UserReq struct {
	MODEL
	UserName string      `gorm:"size:36" json:"user_name"`                  // 用户名
	Role     consts.Role `gorm:"size:4;default:1" json:"role,select(info)"` // 权限  1 管理员  2 普通用户  3 游客
}
