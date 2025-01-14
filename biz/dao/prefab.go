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

// PGetPrefabs 分页获取prefab
func PGetPrefabs(ctx context.Context, pageNum int, pageSize int) ([]*model.Prefab, int64, error) {
	var prefabs []*model.Prefab
	var total int64
	err := infra.MysqlDB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&prefabs).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_get_prefabs_error")
		return nil, 0, err
	}

	err = infra.MysqlDB.Model(&model.Prefab{}).Count(&total).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_get_prefabs_count_error")
		return nil, 0, err
	}
	return prefabs, total, nil
}

// CreatePrefab 在数据库中创建一条prefab记录
func CreatePrefab(ctx context.Context, prefabID int64, prefabName, prefabURL string) error {
	prefab := &model.Prefab{
		PrefabID:   prefabID,
		PrefabName: prefabName,
		PrefabURL:  prefabURL,
	}
	err := infra.MysqlDB.WithContext(ctx).Create(prefab).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_create_prefab_error", "prefabID", prefabID)
		return err
	}
	return nil
}

// UpdatePrefab 更新prefab
func UpdatePrefab(ctx context.Context, prefabID int64, prefabName, prefabURL string) error {
	prefab := &model.Prefab{
		PrefabID:   prefabID,
		PrefabName: prefabName,
		PrefabURL:  prefabURL,
	}
	err := infra.MysqlDB.WithContext(ctx).Model(&model.Prefab{}).Where("prefab_id = ?", prefabID).Updates(prefab).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_update_prefab_error", "prefabID", prefabID)
		return err
	}
	return nil
}

// DeletePrefab 删除prefab
func DeletePrefab(ctx context.Context, prefabID int64) error {
	err := infra.MysqlDB.WithContext(ctx).Where("prefab_id = ?", prefabID).Delete(&model.Prefab{}).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_delete_prefab_error", "prefabID", prefabID)
		return err
	}
	return nil
}
