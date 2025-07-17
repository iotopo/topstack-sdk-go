package asset

import (
	"net/http"
	"topstack-sdk-go/client"
)

// 告警工单列表
// GET /asset/open_api/v1/alert_work_order

type AlertWorkOrderListRequest struct {
	Start      string `json:"start,omitempty"`
	End        string `json:"end,omitempty"`
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DeviceCode string `json:"deviceCode,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
}

func (AlertWorkOrderListRequest) Method() string { return http.MethodGet }
func (AlertWorkOrderListRequest) Url() string    { return "/asset/open_api/v1/alert_work_order" }

type AlertWorkOrderListItem struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	DeviceID    string `json:"deviceID"`
	DeviceName  string `json:"deviceName"`
	DeviceCode  string `json:"deviceCode"`
	Status      int    `json:"status"`
	Content     string `json:"content"`
	CreatedAt   string `json:"createdAt"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	ExecutorID  string `json:"executorID"`
	ExecutorName string `json:"executorName"`
	ReviewerID  string `json:"reviewerID"`
	ReviewerName string `json:"reviewerName"`
	AttachmentExt string `json:"attachmentExt,omitempty"`
}

type AlertWorkOrderListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int                      `json:"total"`
		Items []AlertWorkOrderListItem `json:"items"`
	} `json:"data"`
}

func QueryAlertWorkOrderList(req AlertWorkOrderListRequest) (resp AlertWorkOrderListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警工单详情
// GET /asset/open_api/v1/alert_work_order/:id

type AlertWorkOrderDetailRequest struct {
	ID string
}

func (r AlertWorkOrderDetailRequest) Method() string { return http.MethodGet }
func (r AlertWorkOrderDetailRequest) Url() string    { return "/asset/open_api/v1/alert_work_order/" + r.ID }

type AlertWorkOrderDetailResponse struct {
	Success bool                  `json:"success"`
	Data    AlertWorkOrderListItem `json:"data"`
}

func QueryAlertWorkOrderDetail(req AlertWorkOrderDetailRequest) (resp AlertWorkOrderDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 现场工单列表
// GET /asset/open_api/v1/locale_work_order

type LocaleWorkOrderListRequest struct {
	Start      string `json:"start,omitempty"`
	End        string `json:"end,omitempty"`
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DeviceCode string `json:"deviceCode,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
}

func (LocaleWorkOrderListRequest) Method() string { return http.MethodGet }
func (LocaleWorkOrderListRequest) Url() string    { return "/asset/open_api/v1/locale_work_order" }

type LocaleWorkOrderListItem = AlertWorkOrderListItem

type LocaleWorkOrderListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int                      `json:"total"`
		Items []LocaleWorkOrderListItem `json:"items"`
	} `json:"data"`
}

func QueryLocaleWorkOrderList(req LocaleWorkOrderListRequest) (resp LocaleWorkOrderListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 现场工单详情
// GET /asset/open_api/v1/locale_work_order/:id

type LocaleWorkOrderDetailRequest struct {
	ID string
}

func (r LocaleWorkOrderDetailRequest) Method() string { return http.MethodGet }
func (r LocaleWorkOrderDetailRequest) Url() string    { return "/asset/open_api/v1/locale_work_order/" + r.ID }

type LocaleWorkOrderDetailResponse struct {
	Success bool                  `json:"success"`
	Data    LocaleWorkOrderListItem `json:"data"`
}

func QueryLocaleWorkOrderDetail(req LocaleWorkOrderDetailRequest) (resp LocaleWorkOrderDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 维护工单列表
// GET /asset/open_api/v1/maintenance_work_order

type MaintenanceWorkOrderListRequest struct {
	Start      string `json:"start,omitempty"`
	End        string `json:"end,omitempty"`
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DeviceCode string `json:"deviceCode,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
}

func (MaintenanceWorkOrderListRequest) Method() string { return http.MethodGet }
func (MaintenanceWorkOrderListRequest) Url() string    { return "/asset/open_api/v1/maintenance_work_order" }

type MaintenanceWorkOrderListItem = AlertWorkOrderListItem

type MaintenanceWorkOrderListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int                          `json:"total"`
		Items []MaintenanceWorkOrderListItem `json:"items"`
	} `json:"data"`
}

func QueryMaintenanceWorkOrderList(req MaintenanceWorkOrderListRequest) (resp MaintenanceWorkOrderListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 维护工单详情
// GET /asset/open_api/v1/maintenance_work_order/:id

type MaintenanceWorkOrderDetailRequest struct {
	ID string
}

func (r MaintenanceWorkOrderDetailRequest) Method() string { return http.MethodGet }
func (r MaintenanceWorkOrderDetailRequest) Url() string    { return "/asset/open_api/v1/maintenance_work_order/" + r.ID }

type MaintenanceWorkOrderDetailResponse struct {
	Success bool                      `json:"success"`
	Data    MaintenanceWorkOrderListItem `json:"data"`
}

func QueryMaintenanceWorkOrderDetail(req MaintenanceWorkOrderDetailRequest) (resp MaintenanceWorkOrderDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 计划工单列表
// GET /asset/open_api/v1/schedule_work_order

type ScheduleWorkOrderListRequest struct {
	Start      string `json:"start,omitempty"`
	End        string `json:"end,omitempty"`
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DeviceCode string `json:"deviceCode,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
}

func (ScheduleWorkOrderListRequest) Method() string { return http.MethodGet }
func (ScheduleWorkOrderListRequest) Url() string    { return "/asset/open_api/v1/schedule_work_order" }

type ScheduleWorkOrderListItem = AlertWorkOrderListItem

type ScheduleWorkOrderListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int                      `json:"total"`
		Items []ScheduleWorkOrderListItem `json:"items"`
	} `json:"data"`
}

func QueryScheduleWorkOrderList(req ScheduleWorkOrderListRequest) (resp ScheduleWorkOrderListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 计划工单详情
// GET /asset/open_api/v1/schedule_work_order/:id

type ScheduleWorkOrderDetailRequest struct {
	ID string
}

func (r ScheduleWorkOrderDetailRequest) Method() string { return http.MethodGet }
func (r ScheduleWorkOrderDetailRequest) Url() string    { return "/asset/open_api/v1/schedule_work_order/" + r.ID }

type ScheduleWorkOrderDetailResponse struct {
	Success bool                  `json:"success"`
	Data    ScheduleWorkOrderListItem `json:"data"`
}

func QueryScheduleWorkOrderDetail(req ScheduleWorkOrderDetailRequest) (resp ScheduleWorkOrderDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
} 