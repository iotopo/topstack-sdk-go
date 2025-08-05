package nats

import (
	"time"
)

type PointData struct {
	DeviceID  string      `json:"deviceID"`
	PointID   string      `json:"pointID"`
	Value     interface{} `json:"value"`
	Quality   uint16      `json:"quality"`             // 1 表示离线，2 表示无效
	Timestamp time.Time   `json:"timestamp,omitempty"` // 时间戳，单位:毫秒
	Status    int         `json:"status,omitempty"`    // TODO: 状态: 0 表示正常，> 0 表示越上限，< 0 表示越下限

	// 接收到数据后补充, 然后再转发到 nats 消息总线
	DeviceTypeID string `json:"deviceTypeID"`
	ProjectID    string `json:"projectID"`
	GatewayID    string `json:"gatewayID"`

	NotSave bool `json:"notSave,omitempty"` // 不持久化存储
}

type GatewayState struct {
	SN        string    `json:"sn,omitempty"`
	Name      string    `json:"name,omitempty"`
	ProjectID string    `json:"projectID,omitempty"`
	GatewayID string    `json:"gatewayID,omitempty"`
	State     int       `json:"state"`               // 0: 离线， 1：在线
	Timestamp time.Time `json:"timestamp,omitempty"` // 时间戳
}

// DeviceState 设备状态(在线/离线)
type DeviceState struct {
	ProjectID string    `json:"projectID,omitempty"`
	GatewayID string    `json:"gatewayID,omitempty"`
	DeviceID  string    `json:"deviceID,omitempty"`
	State     int       `json:"state"`               // 0: 离线， 1：在线
	Timestamp time.Time `json:"timestamp,omitempty"` // 时间戳
}

// ChannelState 数据通道状态
type ChannelState struct {
	ProjectID   string    `json:"projectID"`
	GatewayID   string    `json:"gatewayID"`
	ChannelID   string    `json:"channelID"`
	Running     bool      `json:"running"`
	Connected   bool      `json:"connected"`
	Timestamp   time.Time `json:"timestamp"`
	GatewayName string    `json:"gatewayName"`
	ChannelName string    `json:"channelName"`
}

type AlertBase struct {
	ID           string     `json:"id" gorm:"size:36;primaryKey"`
	Status       string     `json:"status" gorm:"size:10;index"` // unhandled/handled/ignored/auto, 默认为 unhandled
	CreatedAt    time.Time  `json:"createdAt"`                   // 创建时间
	RecoveredAt  *time.Time `json:"recoveredAt"`                 // 解除时间
	HandledAt    *time.Time `json:"handledAt"`                   // 确认或忽略时间
	ExpiredAt    *time.Time `json:"expiredAt"`                   // 过期时间
	Handler      string     `json:"handler" gorm:"size:64"`      // 确认人或忽略人
	OrderCreated bool       `json:"orderCreated"`                // 是否创建过工单
	Edge         bool       `json:"edge,omitempty"`              // 是否为边缘侧产生
	Title        string     `json:"title" gorm:"size:50"`        // 通知标题
	Content      string     `json:"content" gorm:"type:text"`    // 通知内容
	Remark       string     `json:"remark" gorm:"type:text"`     // 确认信息的备注

	RuleTemplateID *string `json:"ruleTemplateID" gorm:"size:36"`

	// 触发器信息
	TriggerID    string  `json:"triggerID" gorm:"size:36"`
	Mode         string  `json:"mode" gorm:"size:20"`                  // 触发方式: property(属性值触发), status(上下线触发), logic
	CompareMode  string  `json:"compareMode" gorm:"size:20"`           // 比较模式
	CompareValue string  `json:"compareValue" gorm:"size:255"`         // 比较值
	Duration     int     `json:"duration,omitempty"`                   // 持续时长（秒）
	InputValue   string  `json:"inputValue,omitempty" gorm:"size:255"` // 输入值
	PointID      string  `json:"pointID" gorm:"size:36"`               // 测点ID匹配
	DeadBand     float64 `json:"deadBand,omitempty"`                   // 死区值
	Diff         float64 `json:"diff,omitempty"`                       // 偏差值

	ProjectID    string  `json:"projectID" gorm:"index;size:36"`
	DeviceID     *string `json:"deviceID" gorm:"index;size:36"`     // 关联
	AlertTypeID  string  `json:"alertTypeID" gorm:"size:36"`        // 告警类型ID
	AlertLevelID string  `json:"alertLevelID" gorm:"size:36;index"` // 告警等级 ID
}

type AlertInfo struct {
	AlertBase
	// 以下字段仅在发生告警时生效
	RuleName        string                 `json:"ruleName"`
	AlertTypeName   string                 `json:"alertTypeName"`
	AlertTypeCode   string                 `json:"alertTypeCode"`
	AlertLevelCode  string                 `json:"alertLevelCode"`
	AlertLevelColor string                 `json:"alertLevelColor"`
	AlertLevelName  string                 `json:"alertLevelName"`
	DeviceName      string                 `json:"deviceName,omitempty"`
	PointName       string                 `json:"pointName,omitempty"`
	DeviceTypeID    string                 `json:"deviceTypeID"`
	DeviceGroupID   string                 `json:"deviceGroupID"`
	DeviceAttr      map[string]interface{} `json:"deviceAttr,omitempty" gorm:"-"`
}
