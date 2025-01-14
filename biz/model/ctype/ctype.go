package ctype

import "fmt"

type VisionType int64 // 预制体类型 1 视频 2 预制体 3 其他

// Scan VisionType Scanner 接口实现
func (v *VisionType) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("VisionType value is nil")
	}
	var str string

	switch value.(type) {
	case string:
		str = value.(string)
	case []byte:
		str = string(value.([]byte))
	}

	switch str {
	case "视频":
		*v = VisionType(1)
	case "预制体":
		*v = VisionType(2)
	default:
		*v = VisionType(3)
	}
	return nil
}

// Value VisionType Valuer 接口实现
func (v VisionType) Value() (string, error) {
	switch v {
	case 1:
		return "视频", nil
	case 2:
		return "预制体", nil
	default:
		return "其他", nil
	}
}
