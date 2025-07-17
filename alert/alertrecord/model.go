package alertrecord

type AlertRecordActivityItem struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	DeviceID  string `json:"deviceID"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type ActivityRequest struct {
	Start         string `json:"start,omitempty"`
	End           string `json:"end,omitempty"`
	AlertTypeID   string `json:"alertTypeID,omitempty"`
	DeviceGroupID string `json:"deviceGroupID,omitempty"`
	DeviceID      string `json:"deviceID,omitempty"`
	DeviceTags    string `json:"deviceTags,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

type ActivityResponse struct {
	Total int                       `json:"total"`
	Items []AlertRecordActivityItem `json:"items"`
}

type IgnoredBatchRequest struct {
	IDs []string `json:"ids"`
}

type IgnoredBatchResponse struct{} 