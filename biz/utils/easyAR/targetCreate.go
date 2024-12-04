package easyAR

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

// signedRequest 表示带签名的参数
type signedRequest struct {
	Request   interface{}
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
func CreateTarget(request TargetRequest) (string, error) {
	params := &TargetRequest{
		Name:  request.Name,
		Image: request.Image,
		Type:  request.Type,
		Size:  request.Size,
		Meta:  request.Meta,
	}

	// 这里的 signedParams 是一个map[string]interface{}
	signedParams, err := signParam(params)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(signedParams)
	fmt.Println(string(jsonData))
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
	fmt.Println(string(respBody))
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

// @Title signParam
// @Description 对前端传来的参数进行签名
// @Param params 接口类型，其动态类型是request指针
func signParam(params interface{}) (map[string]interface{}, error) {
	timestamp := time.Now().UnixNano() / 1e6
	easyARConf := conf.Conf.EasyAR

	var paramMap map[string]interface{}

	// 通过类型断言判断 `params` 的动态类型
	switch v := params.(type) {
	case *TargetRequest:
		// 如果是 TargetRequest 指针类型
		paramMap = map[string]interface{}{
			"name":      v.Name,
			"image":     v.Image,
			"type":      v.Type,
			"size":      v.Size,
			"meta":      v.Meta,
			"timestamp": timestamp,
			"appId":     easyARConf.CrsAppID,
			"apiKey":    easyARConf.ApiKey,
		}
	case *TargetListRequest:
		// 如果是 ListTargetRequest 指针类型
		paramMap = map[string]interface{}{
			"pageNum":   v.PageNum,
			"pageSize":  v.PageSize,
			"timestamp": timestamp,
			"appId":     easyARConf.CrsAppID,
			"apiKey":    easyARConf.ApiKey,
		}
	case *TargetInfoRequest:
		// 如果是 TargetInfoRequest 指针类型
		paramMap = map[string]interface{}{
			"timestamp": timestamp,
			"appId":     easyARConf.CrsAppID,
			"apiKey":    easyARConf.ApiKey,
		}
	default:
		return nil, fmt.Errorf("unsupported parameter type: %T", params)
	}

	keys := make([]string, 0, len(paramMap))
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	paramStr := strings.Join(sortKeysAndValues(keys, paramMap), "")
	signature := generateSignature(paramStr, easyARConf.ApiSecret)

	paramMap["signature"] = signature

	return paramMap, nil
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
