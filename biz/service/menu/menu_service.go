package menu_service

import (
	"context"
	"uav/biz/dao"
	"uav/biz/model"
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
