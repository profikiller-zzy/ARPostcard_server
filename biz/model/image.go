package model

// Image 识别图
type Image struct {
	MODEL
	ImageID   string `gorm:"size:36" json:"image_id"`   // 图片ID，这个是EasyAR返回的图片ID
	ImageURL  string `gorm:"size:128" json:"image_url"` // 图片URL，存储到对象存储服务当中的URL
	ImageName string `gorm:"size:64" json:"image_name"` // 图片名称
	PrefabID  int64  `gorm:"column:prefab_id" json:"prefab_id"`
	VideoID   int64  `gorm:"column:video_id" json:"video_id"`
}

func (Image) TableName() string {
	return "images"
}
