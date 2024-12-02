package handler

import (
	image_serveice "ARPostcard_server/biz/service/image"
	"ARPostcard_server/biz/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func ImageCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_serveice.TargetRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	err = image_serveice.ImageCreate(ctx, req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, nil)
}

func GetPrefabName(ctx context.Context, c *app.RequestContext) {
	var err error
	var req image_serveice.PrefabNameRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	prefabName, err := image_serveice.GetPrefabName(ctx, req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, prefabName)
}
