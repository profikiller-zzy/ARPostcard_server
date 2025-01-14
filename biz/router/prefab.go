package router

import (
	"ARPostcard_server/biz/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func PrefabRegister(r route.IRouter) {
	_prefab := r.Group("/prefab")
	{
		_prefab.GET("/prefabs", handler.PrefabList)
	}
}
