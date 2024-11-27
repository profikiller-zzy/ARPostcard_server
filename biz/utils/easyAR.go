package utils

import (
	"ARPostcard_server/biz/conf"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type TargetRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Type  string `json:"type"`
	Size  string `json:"size"`
	Meta  string `json:"meta"`
}

// SignedParams 表示带签名的参数
type signedParams struct {
	*TargetRequest
	Timestamp int64  `json:"timestamp"`
	AppId     string `json:"appId"`
	ApiKey    string `json:"apiKey"`
	Signature string `json:"signature"`
}

// TargetResponse 表示目标上传的响应
type TargetResponse struct {
	StatusCode int          `json:"statusCode"`
	Result     TargetResult `json:"result"`
	Timestamp  int64        `json:"timestamp"`
}

// TargetResult 包含目标ID
type TargetResult struct {
	TargetId string `json:"targetId"`
}

// CreateTarget 创建新的目标
func CreateTarget(imgInfo TargetRequest) (string, error) {
	params := &TargetRequest{
		Name:  imgInfo.Name,
		Image: imgInfo.Image,
		Type:  imgInfo.Type,
		Size:  imgInfo.Size,
		Meta:  imgInfo.Meta,
	}

	signedParams, err := signParam(params)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(signedParams)
	if err != nil {
		return "", err
	}

	cloudURL := conf.Conf.EasyAR.CloudURL
	req, err := http.NewRequest("POST", cloudURL+"/targets", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var targetResponse TargetResponse
	if err := json.Unmarshal(respBody, &targetResponse); err != nil {
		return "", err
	}
	if targetResponse.StatusCode != 0 {
		return "", fmt.Errorf("EasyAR API returned an error: %v", targetResponse)
	}

	return targetResponse.Result.TargetId, nil
}

// signParam 对请求参数进行签名
func signParam(params *TargetRequest) (*signedParams, error) {
	timestamp := time.Now().UnixNano() / 1e6
	easyARConf := conf.Conf.EasyAR
	paramMap := map[string]interface{}{
		"name":      params.Name,
		"image":     params.Image,
		"type":      params.Type,
		"size":      params.Size,
		"meta":      params.Meta,
		"timestamp": timestamp,
		"appId":     easyARConf.CrsAppID,
		"apiKey":    easyARConf.ApiKey,
	}

	keys := make([]string, 0, len(paramMap))
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	paramStr := strings.Join(sortKeysAndValues(keys, paramMap), "")
	signature := generateSignature(paramStr, easyARConf.ApiSecret)

	signedParams := &signedParams{
		TargetRequest: params,
		Timestamp:     timestamp,
		AppId:         easyARConf.CrsAppID,
		ApiKey:        easyARConf.ApiKey,
		Signature:     signature,
	}

	return signedParams, nil
}

// sortKeysAndValues 对参数键值对进行排序并拼接
func sortKeysAndValues(keys []string, params map[string]interface{}) []string {
	var parts []string
	for _, key := range keys {
		value := params[key]
		parts = append(parts, key+fmt.Sprintf("%v", value))
	}
	return parts
}

// generateSignature 使用SHA256生成签名
func generateSignature(paramStr, secret string) string {
	hash := sha256.Sum256([]byte(paramStr + secret))
	return fmt.Sprintf("%x", hash)
}
