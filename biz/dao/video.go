package dao

import (
	"ARPostcard_server/biz/consts"
	"ARPostcard_server/biz/infra"
	"ARPostcard_server/biz/model"
	"ARPostcard_server/biz/utils"
	"context"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
)

func CreateVideo(ctx context.Context, videoName, videoURL string) (int64, error) {
	videoId := utils.RandomInt64()
	video := &model.Video{
		VideoID:   videoId,
		VideoURL:  videoURL,
		VideoName: videoName,
	}

	err := infra.MysqlDB.WithContext(ctx).Debug().
		Create(video).Error

	if err != nil {
		ilog.EventError(ctx, err, "dao_create_video_error", "videoUrl", videoURL)
		return -1, ierror.NewIError(consts.DBError, err.Error())
	}

	return videoId, nil
}
