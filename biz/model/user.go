package model

import "uav/biz/consts"

// User 用户表
type User struct {
	MODEL
	NickName string      `gorm:"size:36" json:"nick_name,select(c|info)"`   // 昵称
	UserName string      `gorm:"size:36" json:"user_name"`                  // 用户名
	Password string      `gorm:"size:128" json:"password"`                  // 密码
	Avatar   string      `gorm:"size:256" json:"avatar,select(c)"`          // 头像
	Email    string      `gorm:"size:128" json:"email,select(info)"`        // 邮箱
	Tel      string      `gorm:"size:18" json:"tel"`                        // 手机号
	Addr     string      `gorm:"size:64" json:"addr,select(c|info)"`        // 地址
	IP       string      `gorm:"size:20" json:"ip,select(c)"`               // ip地址
	Role     consts.Role `gorm:"size:4;default:1" json:"role,select(info)"` // 权限  1 管理员  2 普通用户  3 游客
}
