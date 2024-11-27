package image_serveice

import (
	"ARPostcard_server/biz/dao"
	"ARPostcard_server/biz/utils"
	"context"
)

// TargetRequest 表示目标上传的请求
type TargetRequest struct {
	utils.TargetRequest
	PrefabName string `json:"prefab_name"`
}

func ImageCreate(ctx context.Context, req TargetRequest) error {
	imageID, err := utils.CreateTarget(req.TargetRequest)
	if err != nil {
		return err
	}

	err = dao.CreateImage(ctx, imageID, "", req.PrefabName)
	if err != nil {
		return err
	}

	return nil
}
