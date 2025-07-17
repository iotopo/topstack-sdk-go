package locale

type ListRequest struct {
	Start      string `json:"start,omitempty"`
	End        string `json:"end,omitempty"`
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DeviceCode string `json:"deviceCode,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
}

type ListItem struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	DeviceID      string `json:"deviceID"`
	DeviceName    string `json:"deviceName"`
	DeviceCode    string `json:"deviceCode"`
	Status        int    `json:"status"`
	Content       string `json:"content"`
	CreatedAt     string `json:"createdAt"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	ExecutorID    string `json:"executorID"`
	ExecutorName  string `json:"executorName"`
	ReviewerID    string `json:"reviewerID"`
	ReviewerName  string `json:"reviewerName"`
	AttachmentExt string `json:"attachmentExt,omitempty"`
}

type ListResponse struct {
	Total int       `json:"total"`
	Items []ListItem `json:"items"`
}

type DetailResponse ListItem 