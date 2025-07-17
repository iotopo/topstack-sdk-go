package iot

import (
	"net/http"
	"time"
)

type DeviceTypeQueryRequest struct {
	Search    string `json:"search,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	PageNum   int    `json:"pageNum,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

func (DeviceTypeQueryRequest) Method() string {
	return http.MethodGet
}
func (DeviceTypeQueryRequest) Url() string {
	return "/iot/open_api/v1/device_type/query"
}

type DeviceTypeQueryResponse struct {
	Total int64 `json:"total"`
	Items []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Group       string `json:"group"` // 分组
		Code        string `json:"code"`
		Description string `json:"description"`

		CreatedAt time.Time `json:"createdAt,omitempty"` // 创建时间
		UpdatedAt time.Time `json:"updatedAt,omitempty"` // 更新时间

		LongitudePointID string `json:"longitudePointID"` // 经度（地理坐标）点位
		LatitudePointID  string `json:"latitudePointID"`  // 纬度（地理坐标）点位
	} `json:"items"`
}
