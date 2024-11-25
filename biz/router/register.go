package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"uav/biz/mw/response_header"
)

func Register(r *server.Hertz) {
	uav := r.Group("/uav")
	uav.Use(response_header.RespLog())

	MenuRegister(uav)
	UserRegister(uav)
}
