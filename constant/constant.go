package constant

const (
	SCHEME = "https://"
	HOST   = "pq2r6vuwuh.re.qweatherapi.com"
	APIKEY = "caa94a8d47ba4daba3f24f11e9aeaee5"
)

// Refer 引用信息（如果已经在其他文件定义过，可以删除这个结构体）
type Refer struct {
	Sources []string `json:"sources"`
	License []string `json:"license"`
}

// const (
//
//	INVALID PARAMETER
//
// )
// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Error struct {
		Status int    `json:"status"`
		Type   string `json:"type"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
	} `json:"error"`
}
