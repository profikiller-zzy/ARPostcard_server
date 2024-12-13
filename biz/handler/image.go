package handler

import (
	image_service "ARPostcard_server/biz/service/image"
	"ARPostcard_server/biz/utils"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

func ImageCreate(ctx context.Context, c *app.RequestContext) {
	var err error

	req, err := image_service.GetImageInfoFromForm(ctx, c)

	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	err = image_service.ImageCreate(ctx, *req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, nil)
}

func GetVideoName(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_service.VideoNameRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	videoName, err := image_service.GetVideoName(ctx, req)
	fmt.Println(videoName)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, videoName)
}

func GetPrefabName(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_service.PrefabNameRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	prefabName, err := image_service.GetPrefabName(ctx, req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, prefabName)
}

func GetImageList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_service.TargetListRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	imageIDs, err := image_service.GetImageList(ctx, req)

	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, imageIDs)
}

// GetImageListFromDB 从数据库当中查询当前的图片，主要用于后台管理
func GetImageListFromDB(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_service.TargetListRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	images, total, err := image_service.GetImageListFromDB(ctx, req)

	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	data := map[string]interface{}{
		"list":  images,
		"total": total,
	}
	utils.RespOK(ctx, c, data)
}

// GetImageInfo 获取图片的详细信息，先获取数据库当中的信息，然后获取在easyAR云端的信息
func GetImageInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_service.ImageInfoRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	image, err := image_service.GetImageInfo(ctx, req)

	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, image)
}

// DeleteImage 删除图片，同时删除数据库当中的信息和easyAR云端的信息，可以采用设置的标志位来实现，但是开始的时候先不考虑
func DeleteImage(ctx context.Context, c *app.RequestContext) {

}

// UpdateImage 更新图片，可以更新数据库当中的信息和easyAR云端的信息，可以采用设置的标志位来实现，但是开始的时候先不考虑
func UpdateImage(ctx context.Context, c *app.RequestContext) {
}
