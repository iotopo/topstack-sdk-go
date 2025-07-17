package devicetype

import (
	"time"
)

// Device Type related models
type QueryRequest struct {
	Search    string `json:"search,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	PageNum   int    `json:"pageNum,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type QueryResponse struct {
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

// Device Type Point related models
type PointQueryRequest struct {
	Search       string `json:"search,omitempty"`
	DeviceTypeID string `json:"deviceTypeID,omitempty"`
	Type         string `json:"type,omitempty"`
	Order        string `json:"order,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

type PointQueryResponse struct {
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