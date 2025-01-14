package service

type PListRequest struct {
	PageNum  int `json:"pageNum" query:"pageNum"`
	PageSize int `json:"pageSize" query:"pageSize"`
}
