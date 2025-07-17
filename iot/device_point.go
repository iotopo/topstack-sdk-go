package iot

import "net/http"

type DevicePointQueryRequest struct {
	Search   string `json:"search,omitempty"`
	DeviceID string `json:"deviceID,omitempty"`
	Type     string `json:"type,omitempty"`
	Order    string `json:"order,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

func (DevicePointQueryRequest) Method() string {
	return http.MethodGet
}

func (DevicePointQueryRequest) Url() string {
	return "/iot/open_api/v1/device_point/query"
}

type DevicePointQueryResponse struct {
	Total int64 `json:"total"`
	Items []struct {
		PointID    string `json:"pointID"`
		Name       string `json:"name"`
		Type       string `json:"type"`       // 类型 int double string bool array float time
		AccessMode string `json:"accessMode"` // 读写类型，只读(r)、只写(w)、读写(rw)

		OrderNumber int    `json:"orderNumber"`
		Description string `json:"description,omitempty"`
		Group       string `json:"group,omitempty"`
		Unit        string `json:"unit,omitempty"`      // 计量单位
		Format      string `json:"format,omitempty"`    // 格式化 0.00 http://numeraljs.com/#format
		Edge        bool   `json:"edge,omitempty"`      // 是否只在边缘侧使用
		IsArray     bool   `json:"isArray,omitempty"`   // 是否为数组
		Generator   string `json:"generator,omitempty"` // 模拟器函数, 当通道启用仿真模式时生效
	} `json:"items"`
}
