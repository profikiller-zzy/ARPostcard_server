package router

import (
	"github.com/cloudwego/hertz/pkg/route"
	"uav/biz/handler"
	"uav/biz/mw/jwt_auth"
)

func MenuRegister(r route.IRouter) {
	_menu := r.Group("/menu")
	{
		_menu.GET("", jwt_auth.JwtAuth(), handler.MenuList)
	}
}
