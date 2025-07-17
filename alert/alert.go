package alert

import (
	"net/http"
	"topstack-sdk-go/client"
)

// 告警等级查询
// GET /alert/open_api/v1/alert_level

type AlertLevelQueryRequest struct {
	// 可以根据需要添加查询参数
}

func (AlertLevelQueryRequest) Method() string { return http.MethodGet }
func (AlertLevelQueryRequest) Url() string    { return "/alert/open_api/v1/alert_level" }

type AlertLevelItem struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
	Color string `json:"color"`
	Label string `json:"label"`
}

type AlertLevelQueryResponse struct {
	Success bool           `json:"success"`
	Data    []AlertLevelItem `json:"data"`
}

func QueryAlertLevel() (resp AlertLevelQueryResponse, err error) {
	err = client.SendRequest(AlertLevelQueryRequest{}, &resp)
	return
}

// 告警等级新增
// POST /alert/open_api/v1/alert_level

type AlertLevelCreateRequest struct {
	Value int    `json:"value"`
	Color string `json:"color"`
	Label string `json:"label"`
}

func (AlertLevelCreateRequest) Method() string { return http.MethodPost }
func (AlertLevelCreateRequest) Url() string    { return "/alert/open_api/v1/alert_level" }

type AlertLevelCreateResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func CreateAlertLevel(req AlertLevelCreateRequest) (resp AlertLevelCreateResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警等级修改
// PUT /alert/open_api/v1/alert_level

type AlertLevelUpdateRequest struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
	Color string `json:"color"`
	Label string `json:"label"`
}

func (AlertLevelUpdateRequest) Method() string { return http.MethodPut }
func (AlertLevelUpdateRequest) Url() string    { return "/alert/open_api/v1/alert_level" }

type AlertLevelUpdateResponse = AlertLevelCreateResponse

func UpdateAlertLevel(req AlertLevelUpdateRequest) (resp AlertLevelUpdateResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警等级删除
// DELETE /alert/open_api/v1/alert_level

type AlertLevelDeleteRequest struct {
	ID string `json:"id"`
}

func (AlertLevelDeleteRequest) Method() string { return http.MethodDelete }
func (AlertLevelDeleteRequest) Url() string    { return "/alert/open_api/v1/alert_level" }

type AlertLevelDeleteResponse = AlertLevelCreateResponse

func DeleteAlertLevel(req AlertLevelDeleteRequest) (resp AlertLevelDeleteResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警类型分页查询
// GET /alert/open_api/v1/alert_type

type AlertTypeListRequest struct {
	Name     string `json:"name,omitempty"`
	Level    string `json:"level,omitempty"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

func (AlertTypeListRequest) Method() string { return http.MethodGet }
func (AlertTypeListRequest) Url() string    { return "/alert/open_api/v1/alert_type" }

type AlertTypeListItem struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	AlertLevelID string   `json:"alertLevelID"`
	NotifyChannels []string `json:"notifyChannels"`
	OccurContent  string   `json:"occurContent"`
	RecoverContent string  `json:"recoverContent"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
}

type AlertTypeListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int                `json:"total"`
		Types []AlertTypeListItem `json:"types"`
	} `json:"data"`
}

func QueryAlertTypeList(req AlertTypeListRequest) (resp AlertTypeListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警类型新增
// POST /alert/open_api/v1/alert_type

type AlertTypeCreateRequest struct {
	Name         string   `json:"name"`
	AlertLevelID string   `json:"alertLevelID"`
	NotifyChannels []string `json:"notifyChannels"`
	OccurContent  string   `json:"occurContent"`
	RecoverContent string  `json:"recoverContent"`
}

func (AlertTypeCreateRequest) Method() string { return http.MethodPost }
func (AlertTypeCreateRequest) Url() string    { return "/alert/open_api/v1/alert_type" }

type AlertTypeCreateResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func CreateAlertType(req AlertTypeCreateRequest) (resp AlertTypeCreateResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警类型修改
// PUT /alert/open_api/v1/alert_type

type AlertTypeUpdateRequest struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	AlertLevelID string   `json:"alertLevelID"`
	NotifyChannels []string `json:"notifyChannels"`
	OccurContent  string   `json:"occurContent"`
	RecoverContent string  `json:"recoverContent"`
}

func (AlertTypeUpdateRequest) Method() string { return http.MethodPut }
func (AlertTypeUpdateRequest) Url() string    { return "/alert/open_api/v1/alert_type" }

type AlertTypeUpdateResponse = AlertTypeCreateResponse

func UpdateAlertType(req AlertTypeUpdateRequest) (resp AlertTypeUpdateResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警类型删除
// DELETE /alert/open_api/v1/alert_type/{}

type AlertTypeDeleteRequest struct {
	ID string
}

func (r AlertTypeDeleteRequest) Method() string { return http.MethodDelete }
func (r AlertTypeDeleteRequest) Url() string    { return "/alert/open_api/v1/alert_type/" + r.ID }

type AlertTypeDeleteResponse = AlertTypeCreateResponse

func DeleteAlertType(req AlertTypeDeleteRequest) (resp AlertTypeDeleteResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 告警记录批量忽略
// PUT /alert/open_api/v1/alert_record/ignoredBatch

type AlertRecordIgnoredBatchRequest struct {
	IDs []string `json:"ids"`
}

func (AlertRecordIgnoredBatchRequest) Method() string { return http.MethodPut }
func (AlertRecordIgnoredBatchRequest) Url() string    { return "/alert/open_api/v1/alert_record/ignoredBatch" }

type AlertRecordIgnoredBatchResponse struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

func IgnoreAlertRecordBatch(req AlertRecordIgnoredBatchRequest) (resp AlertRecordIgnoredBatchResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 活动告警查询
// GET /alert/open_api/v1/alert_record/activity

type AlertRecordActivityRequest struct {
	Start         string `json:"start,omitempty"`
	End           string `json:"end,omitempty"`
	AlertTypeID   string `json:"alertTypeID,omitempty"`
	DeviceGroupID string `json:"deviceGroupID,omitempty"`
	DeviceID      string `json:"deviceID,omitempty"`
	DeviceTags    string `json:"deviceTags,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

func (AlertRecordActivityRequest) Method() string { return http.MethodGet }
func (AlertRecordActivityRequest) Url() string    { return "/alert/open_api/v1/alert_record/activity" }

type AlertRecordActivityItem struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	DeviceID  string `json:"deviceID"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type AlertRecordActivityResponse struct {
	Success bool                      `json:"success"`
	Data    struct {
		Total int                      `json:"total"`
		Items []AlertRecordActivityItem `json:"items"`
	} `json:"data"`
}

func QueryAlertRecordActivity(req AlertRecordActivityRequest) (resp AlertRecordActivityResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}
