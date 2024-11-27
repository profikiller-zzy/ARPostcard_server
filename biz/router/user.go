package router

import (
	"ARPostcard_server/biz/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func UserRegister(r route.IRouter) {
	_user := r.Group("/user")
	{
		_user.GET("/login", handler.UserLogin)
		_user.GET("/register", handler.UserRegister)
	}
}
