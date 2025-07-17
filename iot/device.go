package iot

import (
	"fmt"
	"net/http"
	"time"
)

type DeviceQueryRequest struct {
	Search          string `json:"search,omitempty"`          // 名称或标识关键字
	GatewayID       string `json:"gatewayID,omitempty"`       // 所属网关
	TypeID          string `json:"typeID,omitempty"`          // 所属模型
	ConnectMode     string `json:"connectMode,omitempty"`     // 接入方式：direct、gateway、custom
	DataChannelID   string `json:"dataChannelID,omitempty"`   // 所属数据通道ID
	CustomChannelID string `json:"customChannelID,omitempty"` // 所属自定义通道ID
	State           string `json:"state,omitempty"`           // 状态 true 表示在线, false 表示离线
	UserGroupID     string `json:"userGroupID,omitempty"`

	Empty bool `json:"empty,omitempty"` // true 表示查询未关联任何通道或网关的设备

	GroupID  string `json:"groupID,omitempty"`  // 所属设备分组
	PageNum  int    `json:"pageNum,omitempty"`  // 当前页，起始值为 1, 默认为 1
	PageSize int    `json:"pageSize,omitempty"` // 每页数量，默认为 10
}

func (DeviceQueryRequest) Method() string {
	return http.MethodGet
}

func (DeviceQueryRequest) Url() string {
	return "/iot/open_api/v1/device/query"
}

type DeviceQueryResponse struct {
	Total int64 `json:"total"`
	Items []struct {
		ID          string `json:"id"`
		Code        string `json:"code"`
		Name        string `json:"name"`
		Description string `json:"description"`
		GatewayID   string `json:"gatewayID"`
		GatewayName string `json:"gatewayName"`
		TypeID      string `json:"typeID"`
		TypeName    string `json:"typeName"`
		GroupID     string `json:"groupID"`
		GroupName   string `json:"groupName"`
		Template    bool   `json:"template"`
		Address     string `json:"address"`
		IdleTimeout int    `json:"idleTimeout,omitempty"`
		ConnectMode string `json:"connectMode"`

		UserGroupID   string `json:"userGroupID"`
		UserGroupName string `json:"userGroupName"`

		DataChannelID   string `json:"dataChannelID"`
		DataChannelName string `json:"dataChannelName"`

		CustomChannelID   string `json:"customChannelID"`
		CustomChannelName string `json:"customChannelName"`

		State           int        `json:"state"` // 网关状态, 0 表示在线，1 表示离线
		StateChangeTime *time.Time `json:"stateChangeTime"`
		CreatedAt       *time.Time `json:"createdAt,omitempty"` // 创建时间
		UpdatedAt       *time.Time `json:"updatedAt,omitempty"` // 更新时间

		HasProps bool `json:"hasProps"` // 指示当前设备的类型上是否定义有属性
		// Editable bool `json:"editable"` // 指示用户是否拥有当前设备所在的用户组

		ManualGI  bool     `json:"manualGI"`
		Longitude *float64 `json:"longitude"` // 经度（地理坐标）
		Latitude  *float64 `json:"latitude"`  // 纬度（地理坐标）

		LongitudePointID string `json:"longitudePointID"`
		LatitudePointID  string `json:"latitudePointID"`
	} `json:"items"`
}

type DevicePropsQueryRequest struct {
	DeviceID string `json:"-"`
}

func (DevicePropsQueryRequest) Method() string {
	return http.MethodGet
}

func (req DevicePropsQueryRequest) Url() string {
	return fmt.Sprintf("/iot/open_api/v1/device/%s/props", req.DeviceID)
}

type DevicePropsQueryResponse []struct {
	PropertyID   string `json:"id"`
	PropertyType string `json:"type"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Value        string `json:"value,omitempty"`
}
