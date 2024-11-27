package handler

import (
	"ARPostcard_server/biz/mw/jwt_auth"
	menu_service "ARPostcard_server/biz/service/menu"
	"ARPostcard_server/biz/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// MenuList 获取菜单列表
func MenuList(ctx context.Context, c *app.RequestContext) {
	userClaims, err := jwt_auth.GetClaimsFromCtx(ctx, c)

	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}

	menuList, err := menu_service.GetMenuList(ctx, userClaims.Role)
	if err != nil {
		utils.RespErr(ctx, c, err)
		return
	}
	utils.RespOK(ctx, c, menuList)
}
