package router

import (
	"ARPostcard_server/biz/handler"
	"ARPostcard_server/biz/mw/jwt_auth"
	"github.com/cloudwego/hertz/pkg/route"
)

func MenuRegister(r route.IRouter) {
	_menu := r.Group("/menu")
	{
		_menu.GET("", jwt_auth.JwtAuth(), handler.MenuList)
	}
}
