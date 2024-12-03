package image_serveice

import (
	"ARPostcard_server/biz/dao"
	"ARPostcard_server/biz/utils/easyAR"
	"context"
)

// TargetRequest 表示目标上传的请求
type TargetRequest struct {
	easyAR.TargetRequest
	PrefabName string `json:"prefab_name"`
}

// PrefabNameRequest 表示获取预制体名称的请求
type PrefabNameRequest struct {
	TargetID string `json:"image_id" query:"image_id"`
}

type TargetListRequest struct {
	easyAR.TargetListRequest
}

func ImageCreate(ctx context.Context, req TargetRequest) error {
	imageID, err := easyAR.CreateTarget(req.TargetRequest)
	if err != nil {
		return err
	}

	err = dao.CreateImage(ctx, imageID, "", req.PrefabName)
	if err != nil {
		return err
	}

	return nil
}

func GetPrefabName(ctx context.Context, req PrefabNameRequest) (string, error) {
	image, err := dao.GetImageByImageID(ctx, req.TargetID)
	if err != nil {
		return "", err
	}

	return image.PrefabName, nil
}

func GetImageList(ctx context.Context, req TargetListRequest) ([]string, error) {
	imageIDs, err := easyAR.GetTargetList(req.TargetListRequest)
	if err != nil {
		return nil, err
	}

	return imageIDs, nil
}
