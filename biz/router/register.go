package router

import (
	"ARPostcard_server/biz/mw/response_header"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {
	uav := r.Group("/uav")
	uav.Use(response_header.RespLog())

	MenuRegister(uav)
	UserRegister(uav)
}
