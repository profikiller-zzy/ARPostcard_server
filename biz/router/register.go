package router

import (
	"ARPostcard_server/biz/mw/response_header"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {
	arPostcard := r.Group("/ar_postcard")
	arPostcard.Use(response_header.RespLog())

	MenuRegister(arPostcard)
	UserRegister(arPostcard)
	ImageRegister(arPostcard)
}
