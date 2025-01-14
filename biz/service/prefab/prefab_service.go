package prefab_service

import (
	"ARPostcard_server/biz/dao"
	"ARPostcard_server/biz/model"
	"ARPostcard_server/biz/service"
	"context"
)

func GetPrefabList(ctx context.Context, req service.PListRequest) ([]*model.Prefab, int64, error) {
	return dao.PGetPrefabs(ctx, req.PageNum, req.PageSize)
}
