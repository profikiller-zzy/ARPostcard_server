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

// TargetListRequest 图库的目标图像列表请求`
type TargetListRequest struct {
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
}

type TargetListResponse struct {
	StatusCode int        `json:"statusCode"`
	Result     TargetList `json:"result"`
	Timestamp  int64      `json:"timestamp"`
}

// TargetListItem 结构体表示单个目标的信息
type TargetListItem struct {
	TargetId string `json:"targetId"`
	Name     string `json:"name"`
	Size     string `json:"size"`
	Meta     string `json:"meta"`
	Type     string `json:"type"`
	AppKey   string `json:"appKey"`
	Active   string `json:"active"`
	Modified int64  `json:"modified"`
}

// TargetList 表示 TargetListItem 的父级
type TargetList struct {
	Targets []TargetListItem `json:"targets"`
}

// GetTargetList
// @Title GetTargetList
// @Description 向easyAR云识别库发请求，获取目标图像id列表
// @Param request 前端传来的获取目标图片列表请求
// @Return []string 返回 `targetID` 列表
func GetTargetList(request TargetListRequest) ([]string, error) {
	params := &TargetListRequest{
		PageNum:  request.PageNum,
		PageSize: request.PageSize,
	}

	signedParams, err := signParam(params)
	if err != nil {
		return nil, err
	}

	cloudURL := conf.Conf.EasyAR.CloudURL
	//appID := conf.Conf.EasyAR.CrsAppID

	// 请求的参数是直接通过URL的查询参数传递，而非JSON
	// 构造 URL 查询参数
	queryParams := url.Values{}
	queryParams.Add("pageNum", fmt.Sprintf("%d", params.PageNum))
	queryParams.Add("pageSize", fmt.Sprintf("%d", params.PageSize))
	queryParams.Add("timestamp", fmt.Sprintf("%d", signedParams["timestamp"]))
	queryParams.Add("appId", fmt.Sprintf("%s", signedParams["appId"]))
	queryParams.Add("apiKey", fmt.Sprintf("%s", signedParams["apiKey"]))
	queryParams.Add("signature", fmt.Sprintf("%s", signedParams["signature"]))

	requestURL := fmt.Sprintf("%s/targets/infos?%s", cloudURL, queryParams.Encode())

	// 创建 GET 请求
	req, err := http.NewRequest("GET", requestURL, nil) // GET 请求不需要 Body
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

	var targetListResponse TargetListResponse
	if err := json.Unmarshal(respBody, &targetListResponse); err != nil {
		return nil, err
	}

	// 检查返回状态码
	if targetListResponse.StatusCode != 0 {
		return nil, fmt.Errorf("EasyAR API returned an error: %v", targetListResponse)
	}

	// 提取 targetId 列表
	var targetIDs []string
	for _, target := range targetListResponse.Result.Targets {
		targetIDs = append(targetIDs, target.TargetId)
	}

	return targetIDs, nil
}
