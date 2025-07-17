package gateway

import (
	"time"
)

// Gateway related models
type QueryRequest struct {
	Page        int    `json:"page,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	Search      string `json:"search,omitempty"`
	UserGroupID string `json:"userGroupID,omitempty"`
	IsManaged   string `json:"isManaged,omitempty"`
	IsVirtual   string `json:"isVirtual,omitempty"`
}

type QueryResponse struct {
	Total int64 `json:"total"`
	Items []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Type            string `json:"type"`
		Description     string `json:"description"`
		IP              string `json:"ip"`
		Version         string `json:"version"`     // mqtt version
		SoftVersion     string `json:"softVersion"` // 软件版本号
		FirmVersion     string `json:"firmVersion"` // 固件版本号
		CompressionType string `json:"compressionType"`
		SN              string `json:"sn"`            // 序列号
		Manufacturer    string `json:"manufacturer"`  // 生产厂商
		UserGroupID     string `json:"userGroupID"`   // 用户组ID
		UserGroupName   string `json:"userGroupName"` // 用户名称
		IsManaged       bool   `json:"isManaged"`     // 是否为纳管网关
		IsVirtual       bool   `json:"isVirtual"`     // 是否为虚拟网关

		State           int        `json:"state"` // 网关状态, 0 表示在线，1 表示离线
		StateChangeTime *time.Time `json:"stateChangeTime"`
		CreatedAt       *time.Time `json:"createdAt,omitempty"` // 创建时间
		UpdatedAt       *time.Time `json:"updatedAt,omitempty"` // 更新时间
	} `json:"items"`
} 