package model

// Video 识别图片出来的视频
type Video struct {
	MODEL
	VideoID   int64  `gorm:"column:video_id" json:"video_id"` // 配置成唯一
	VideoURL  string `gorm:"size:128" json:"video_url"`       // 视频URL，存储到对象存储服务当中的URL
	VideoName string `gorm:"size:64" json:"video_name"`       // 视频名称，存储在本地的视频链接
}
