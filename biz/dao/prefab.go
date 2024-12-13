package dao

import (
	"ARPostcard_server/biz/infra"
	"ARPostcard_server/biz/model"
	"context"
	"github.com/RanFeng/ilog"
)

func GetPrefabById(ctx context.Context, id int64) (*model.Prefab, error) {
	prefab := &model.Prefab{}
	err := infra.MysqlDB.Where("prefab_id = ?", id).First(prefab).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_user_list_error")
		return nil, err
	}
	return prefab, nil
}
