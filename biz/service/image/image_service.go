package image_service

import (
	"ARPostcard_server/biz/dao"
	"ARPostcard_server/biz/model"
	"ARPostcard_server/biz/utils/easyAR"
	"ARPostcard_server/biz/utils/qiniu"
	"context"
	"github.com/RanFeng/ilog"
	"github.com/cloudwego/hertz/pkg/app"
	"mime/multipart"
	"strconv"
)

// TargetRequest 表示目标上传的请求
type TargetRequest struct {
	easyAR.TargetRequest
	PrefabId int64 `json:"prefab_name"`
	Video    *multipart.FileHeader
}

// PrefabRequest 表示获取预制体名称的请求
type PrefabRequest struct {
	TargetID string `json:"image_id" query:"image_id"`
}

// VideoNameRequest 表示获取绑定的视频名称的请求
type VideoNameRequest struct {
	TargetID string `json:"image_id" query:"image_id"`
}

type TargetListRequest struct {
	easyAR.TargetListRequest
}

// ImageInfoRequest 表示获取图片信息的请求，这里采用id列表，写一个通用接口
type ImageInfoRequest struct {
	TargetIDs []string `json:"image_ids"`
}

type ImageAllInfo struct {
	ImageID     string `json:"image_id"`
	ImageURL    string `json:"image_url"`
	ImageName   string `json:"image_name"` // 图片名称
	PrefabName  string `json:"prefab_name"`
	Size        int    `json:"size"`         // 图片大小，其实是在场景中的大小，以厘米为单位
	ImageBase64 string `json:"image_base64"` // 图片的base64编码
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ImageInfoResponse 表示获取图片信息的响应
type ImageInfoResponse struct {
	Images []*ImageAllInfo `json:"images"`
}

// PrefabAndVideoResponse 表示获取预制体名称和视频名称的响应
type PrefabAndVideoResponse struct {
	Prefab *model.Prefab `json:"prefab"`
	Video  *model.Video  `json:"video"`
}

func GetImageInfoFromForm(ctx context.Context, requestContext *app.RequestContext) (*TargetRequest, error) {
	req := &TargetRequest{}
	// 处理字符串部分
	req.Name = string(requestContext.FormValue("name"))
	req.Image = string(requestContext.FormValue("image"))
	req.Type = string(requestContext.FormValue("type"))
	req.Size = string(requestContext.FormValue("size"))
	req.Meta = string(requestContext.FormValue("meta"))

	// 处理整数部分
	prefabIdStr := string(requestContext.FormValue("prefab_id"))
	prefabId, err := strconv.ParseInt(prefabIdStr, 10, 64)
	if err != nil {
		ilog.EventError(ctx, err, "parse_prefab_id_error", "prefabId", prefabIdStr)
		return nil, err
	}
	req.PrefabId = prefabId

	// 处理文件部分
	video, err := requestContext.FormFile("video_file")
	if err != nil {
		ilog.EventError(ctx, err, "get_video_file_error")
		return nil, err
	}
	req.Video = video

	return req, nil
}

func ImageCreate(ctx context.Context, req TargetRequest) error {
	imageID, err := easyAR.CreateTarget(req.TargetRequest)
	if err != nil {
		return err
	}

	url, name, err := qiniu.UploadFileToQiniu(req.Video)
	if err != nil {
		return err
	}

	videoID, err := dao.CreateVideo(ctx, name, url)
	if err != nil {
		return err
	}
	err = dao.CreateImage(ctx, imageID, "", req.Name, req.PrefabId, videoID)
	if err != nil {
		return err
	}

	return nil
}

func GetPrefabAndVideo(ctx context.Context, req PrefabRequest) (*PrefabAndVideoResponse, error) {
	image, err := dao.GetImageByImageID(ctx, req.TargetID)
	if err != nil {
		return nil, err
	}
	var video *model.Video
	var prefab *model.Prefab

	// 获取prefab
	prefab, err = dao.GetPrefabById(ctx, image.PrefabID)
	if err != nil {
		return nil, err
	}

	if image.PrefabID == 1 {
		// 此时代表是视频
		video, err = dao.GetVideoById(ctx, image.VideoID)
		if err != nil {
			return nil, err
		}
	}

	result := &PrefabAndVideoResponse{
		Video:  video,
		Prefab: prefab,
	}

	return result, nil
}

func GetImageList(ctx context.Context, req TargetListRequest) ([]string, error) {
	imageIDs, err := easyAR.GetTargetList(req.TargetListRequest)
	if err != nil {
		return nil, err
	}

	return imageIDs, nil
}

func GetImageListFromDB(ctx context.Context, req TargetListRequest) ([]*model.Image, int64, error) {
	images, total, err := dao.PGetImages(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, 0, err
	}

	return images, total, nil
}

func GetImageInfo(ctx context.Context, req ImageInfoRequest) (*ImageInfoResponse, error) {
	// 遍历判断每一张图片是否存在，如果查询报错，则给当前图片位置填充空值
	images := make([]*ImageAllInfo, 0)
	for _, imageID := range req.TargetIDs {
		image, err := dao.GetImageByImageID(ctx, imageID)
		var imageInfo *ImageAllInfo
		if err != nil {
			imageInfo = &ImageAllInfo{
				ImageID: imageID,
			}
		} else {
			imageInfo = &ImageAllInfo{
				ImageID:   image.ImageID,
				ImageURL:  image.ImageURL,
				ImageName: image.ImageName,
				CreatedAt: image.CreatedAt,
				UpdatedAt: image.UpdatedAt,
			}
		}

		// 获取EasyAR中的内容
		target, err := easyAR.GetTargetInfo(imageID)
		if err != nil {
			// 打印错误
			ilog.EventError(ctx, err, "easyar_get_image_error", "imageID", imageID)
			imageInfo.Size = 0
			imageInfo.ImageBase64 = ""
		} else {
			imageInfo.Size = target.Result.Size
			imageInfo.ImageBase64 = target.Result.TrackingImage
		}

		images = append(images, imageInfo)
	}

	return &ImageInfoResponse{
		Images: images,
	}, nil

}
