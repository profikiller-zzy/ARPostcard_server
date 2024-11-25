package dao

import (
	"context"
	"uav/biz/infra"
	"uav/biz/model"
)

func GetMenuList(ctx context.Context, role int) ([]model.Menu, error) {
	var menus []model.Menu
	err := infra.MysqlDB.WithContext(ctx).Where("role_level >= ?", role).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
