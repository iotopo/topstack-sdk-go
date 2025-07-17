package alerttype

type AlertTypeListItem struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	AlertLevelID   string   `json:"alertLevelID"`
	NotifyChannels []string `json:"notifyChannels"`
	OccurContent   string   `json:"occurContent"`
	RecoverContent string   `json:"recoverContent"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
}

type ListRequest struct {
	Name     string `json:"name,omitempty"`
	Level    string `json:"level,omitempty"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type ListResponse struct {
	Total int                 `json:"total"`
	Types []AlertTypeListItem `json:"types"`
}

type CreateRequest struct {
	Name           string   `json:"name"`
	AlertLevelID   string   `json:"alertLevelID"`
	NotifyChannels []string `json:"notifyChannels"`
	OccurContent   string   `json:"occurContent"`
	RecoverContent string   `json:"recoverContent"`
}

type CreateResponse struct{}

type UpdateRequest struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	AlertLevelID   string   `json:"alertLevelID"`
	NotifyChannels []string `json:"notifyChannels"`
	OccurContent   string   `json:"occurContent"`
	RecoverContent string   `json:"recoverContent"`
}

type UpdateResponse struct{}

type DeleteRequest struct {
	ID string
}

type DeleteResponse struct{} 