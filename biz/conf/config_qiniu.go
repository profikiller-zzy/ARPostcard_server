package conf

type Qiniu struct {
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
	Bucket    string `yaml:"bucket" json:"bucket"` // 存储桶名称
	CDN       string `yaml:"cdn" json:"cdn"`
	Zone      string `yaml:"zone" json:"zone"` // 存储区域
	Size      int64  `yaml:"size" json:"size"` // Size字段保留，如果后续需要上传图片到七牛云，超过 `size`MB的图片允许上传
}
