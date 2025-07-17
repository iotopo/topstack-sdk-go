package iot

import "net/http"

type DeviceTypePointQueryRequest struct {
	Search       string `json:"search,omitempty"`
	DeviceTypeID string `json:"deviceTypeID,omitempty"`
	Type         string `json:"type,omitempty"`
	Order        string `json:"order,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

func (DeviceTypePointQueryRequest) Method() string {
	return http.MethodGet
}

func (DeviceTypePointQueryRequest) Url() string {
	return "/iot/open_api/v1/device_type_point/query"
}

type DeviceTypePointQueryResponse DevicePointQueryResponse
