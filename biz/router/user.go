package router

import (
	"ARPostcard_server/biz/handler"
	"ARPostcard_server/biz/mw/jwt_auth"
	"github.com/cloudwego/hertz/pkg/route"
)

func UserRegister(r route.IRouter) {
	_user := r.Group("/user")
	{
		_user.POST("/login", handler.UserLogin)
		_user.POST("/register", handler.UserRegister)
		_user.GET("", jwt_auth.JwtAuth(), handler.UserList)
		_user.POST("/create", jwt_auth.JwtAuth(), handler.UserCreate)
		_user.POST("/update", jwt_auth.JwtAuth(), handler.UserUpdate)
		_user.POST("/delete", jwt_auth.JwtAuth(), handler.UserDelete)
	}
}
