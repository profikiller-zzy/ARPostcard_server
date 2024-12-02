package router

import (
	"ARPostcard_server/biz/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func ImageRegister(r route.IRouter) {
	_image := r.Group("/image")
	{
		_image.POST("", handler.ImageCreate)
		_image.GET("/prefab_name", handler.GetPrefabName)
	}
}
