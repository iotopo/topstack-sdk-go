package device_type_group

type QueryRequest struct {
	Search   string `json:"search,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type GroupItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type QueryResponse struct {
	Total int64       `json:"total"`
	Items []GroupItem `json:"items"`
}

type CreateRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description,omitempty"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type ModifyRequest struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description,omitempty"`
}

type DeleteRequest struct {
	ID string `json:"id" binding:"required"`
}

// 批量创建设备类型分组请求
// 直接复用 CreateRequest 结构体数组

type BatchCreateRequest []CreateRequest 