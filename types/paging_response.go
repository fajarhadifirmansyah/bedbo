package types

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
}

type ResponsePaging struct {
	Response
	CurrentPage int64       `json:"current_page"`
	PageSize    int64       `json:"page_size"`
	TotalItems  int64       `json:"total_items"`
	TotalPage   int64       `json:"total_page"`
}
