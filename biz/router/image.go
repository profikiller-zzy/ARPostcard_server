package router

import (
	"ARPostcard_server/biz/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func ImageRegister(r route.IRouter) {
	_image := r.Group("/image")
	{
		_image.POST("", handler.ImageCreate)
		_image.GET("/model_url", handler.GetModel)
		_image.GET("/images", handler.GetImageListFromDB)
		_image.POST("/image_info", handler.GetImageInfo)
		_easyAR := _image.Group("/easyAR")
		_easyAR.GET("/image_list", handler.GetImageList)
	}
}
