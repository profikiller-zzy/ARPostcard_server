package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"uav/biz/handler"
)

func UserRegister(r route.IRouter) {
	_user := r.Group("/user")
	{
		_user.GET("/login", handler.UserLogin)
		_user.GET("/register", handler.UserRegister)
	}
}
