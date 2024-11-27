package model

// Image 识别图
type Image struct {
	MODEL
	ImageID    string `gorm:"size:36" json:"image_id"`     // 图片ID，这个是EasyAR返回的图片ID
	ImageURL   string `gorm:"size:128" json:"image_url"`   // 图片URL，存储到对象存储服务当中的URL
	PrefabName string `gorm:"size:128" json:"prefab_name"` // 对应的模型的名称，unity获取之后，可以加载本地的Prefab
}

func (Image) TableName() string {
	return "images"
}
