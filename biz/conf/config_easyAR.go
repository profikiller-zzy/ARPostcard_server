package conf

type EasyAR struct {
	TargetMgmtURL string `json:"target_mgmt_url" yaml:"target_mgmt_url"` // 目标管理URL
	CrsAppID      string `json:"crs_app_id" yaml:"crs_app_id"`           // CRS App ID
	ApiKey        string `json:"api_key" yaml:"api_key"`                 // API Key
	ApiSecret     string `json:"api_secret" yaml:"api_secret"`           // API Secret
}
