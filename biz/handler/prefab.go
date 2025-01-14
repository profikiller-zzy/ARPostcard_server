package handler

import (
	"ARPostcard_server/biz/service"
	prefab_service "ARPostcard_server/biz/service/prefab"
	"ARPostcard_server/biz/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func PrefabList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.PListRequest

	if err = c.BindAndValidate(&req); err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	result, total, err := prefab_service.GetPrefabList(ctx, req)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	utils.RespOK(ctx, c, map[string]interface{}{
		"list":  result,
		"total": total,
	})
}
