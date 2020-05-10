package entity

import "apimake/cmd/flag"

const (
	TableApi = "table_api"
)

type Api struct {
	Code           string
	Name           string
	Description    string
	HttpUri        string      `json:"http_uri"`
	HttpMethod     string      `json:"http_method"`
	HeaderParams   ApiHeader   `json:"header_params"`
	RequestParams  ApiRequest  `json:"request_params"`
	ResponseParams ApiResponse `json:"response_params"`
}

func (Api) TableName() string {
	return flag.FilePath
}

type ApiHeader struct {
	Items []ApiHeaderParam `json:"items"`
}

type ApiRequest struct {
	Items []ApiRequestParam `json:"items"`
}

type ApiResponse struct {
	Items []ApiResponseParam `json:"items"`
}

type ApiHeaderParam struct {
	Tag        string `json:"tag"`
	IsRequired int    `json:"is_required"`
	Content    string `json:"content"`
}

type ApiRequestParam struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	ParentName   string `json:"parent_name"`
	IsRequired   int    `json:"is_required"`
	ExampleValue string `json:"example_value"`
}

type ApiResponseParam struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	ParentName   string `json:"parent_name"`
	IsRequired   int    `json:"is_required"`
	ExampleValue string `json:"example_value"`
}
