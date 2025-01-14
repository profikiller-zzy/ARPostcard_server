package model

import "ARPostcard_server/biz/model/ctype"

// Image 识别图
type Image struct {
	MODEL
	ImageID    string           `gorm:"size:36" json:"image_id"`   // 图片ID，这个是EasyAR返回的图片ID
	ImageURL   string           `gorm:"size:128" json:"image_url"` // 图片URL，存储到对象存储服务当中的URL
	ImageName  string           `gorm:"size:64" json:"image_name"` // 图片名称
	VisionType ctype.VisionType `gorm:"size:4" json:"vision_type"` // 关联模型类型 1 视频 2 预制体 3 其他
	ModelID    int64            `json:"model_id"`                  // 关联模型ID
	//Prefab     Prefab           `gorm:"foreignKey:ModelID;references:PrefabID" json:"prefab"` // 关联的预制体
	//Video      Video            `gorm:"foreignKey:ModelID;references:VideoID" json:"video"`   // 关联的视频
}

func (Image) TableName() string {
	return "images"
}
