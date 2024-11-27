package menu_service

import (
	"ARPostcard_server/biz/dao"
	"ARPostcard_server/biz/model"
	"context"
)

// GetMenuList 获取菜单列表
func GetMenuList(ctx context.Context, role int) ([]*model.Menu, error) {
	menus, err := dao.GetMenuList(ctx, role)
	if err != nil {
		return nil, err
	}
	result := model.BuildMenuTree(menus)

	return result, nil
}
