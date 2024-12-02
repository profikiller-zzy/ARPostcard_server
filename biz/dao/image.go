package dao

import (
	"ARPostcard_server/biz/consts"
	"ARPostcard_server/biz/infra"
	"ARPostcard_server/biz/model"
	"context"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
)

// CreateImage 在数据库中创建一条image记录
func CreateImage(ctx context.Context, imageID string, imageURL string, prefabName string) error {
	image := &model.Image{
		ImageID:    imageID,
		ImageURL:   imageURL,
		PrefabName: prefabName,
	}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Create(image).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_create_image_error", "imageID", imageID)
		return ierror.NewIError(consts.DBError, err.Error())
	}
	return nil
}

// GetImageByImageID 根据imageID获取image
func GetImageByImageID(ctx context.Context, imageID string) (*model.Image, error) {
	image := &model.Image{}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Where("image_id = ?", imageID).
		First(image).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_get_image_by_image_id_error", "imageID", imageID)
		return nil, ierror.NewIError(consts.DBError, err.Error())
	}
	return image, nil
}
