package easyAR

import (
	"ARPostcard_server/biz/conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type TargetInfoRequest struct {
	TargetID string `json:"targetID" query:"targetID"`
}

type TargetInfo struct {
	TargetID                     string  `json:"targetId"`                     // 目标 ID
	TrackingImage                string  `json:"trackingImage"`                // Base64 编码的图像
	Name                         string  `json:"name"`                         // 名称
	Size                         int     `json:"size,string"`                  // 大小（cm），通过 JSON 字符串反序列化为 float64
	Meta                         string  `json:"meta"`                         // 元数据
	Type                         string  `json:"type"`                         // 类型，例如 "ImageTarget"
	Date                         string  `json:"date"`                         // 日期，ISO 8601 格式
	Active                       int     `json:"active,string"`                // 是否启用，0 或 1，字符串解析为整数
	TrackableRate                float64 `json:"trackableRate"`                // 可追踪率
	DetectableRate               float64 `json:"detectableRate"`               // 可检测率
	DetectableDistinctiveness    float64 `json:"detectableDistinctiveness"`    // 可检测的独特性
	DetectableFeatureCount       int     `json:"detectableFeatureCount"`       // 可检测特征数量
	TrackableDistinctiveness     float64 `json:"trackableDistinctiveness"`     // 可追踪的独特性
	TrackableFeatureCount        int     `json:"trackableFeatureCount"`        // 可追踪特征数量
	TrackableFeatureDistribution float64 `json:"trackableFeatureDistribution"` // 可追踪特征分布
	TrackablePatchContrast       float64 `json:"trackablePatchContrast"`       // 可追踪补丁对比度
	TrackablePatchAmbiguity      float64 `json:"trackablePatchAmbiguity"`      // 可追踪补丁歧义
}

type TargetInfoResponse struct {
	StatusCode int        `json:"statusCode"`
	Result     TargetInfo `json:"result"`
	TimeStamp  int64      `json:"timestamp"`
}

func GetTargetInfo(imageID string) (*TargetInfoResponse, error) {
	params := &TargetInfoRequest{
		TargetID: imageID,
	}
	signedParams, err := signParam(params)
	if err != nil {
		return nil, err
	}
	cloudURL := conf.Conf.EasyAR.CloudURL

	queryParams := url.Values{}
	queryParams.Add("timestamp", fmt.Sprintf("%d", signedParams["timestamp"]))
	queryParams.Add("appId", fmt.Sprintf("%s", signedParams["appId"]))
	queryParams.Add("apiKey", fmt.Sprintf("%s", signedParams["apiKey"]))
	queryParams.Add("signature", fmt.Sprintf("%s", signedParams["signature"]))

	requestURL := fmt.Sprintf("%s/target/%s?%s", cloudURL, params.TargetID, queryParams.Encode())
	// 创建请求GET请求
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	if err != nil {
		return nil, err
	}

	var trgetInfoResponse TargetInfoResponse
	if err := json.Unmarshal(respBody, &trgetInfoResponse); err != nil {
		return nil, err
	}

	if trgetInfoResponse.StatusCode != 0 {
		return nil, fmt.Errorf("EasyAR API returned an error: %v", trgetInfoResponse)
	}

	return &trgetInfoResponse, nil
}
