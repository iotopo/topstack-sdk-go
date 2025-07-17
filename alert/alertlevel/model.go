package alertlevel

type AlertLevelItem struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
	Color string `json:"color"`
	Label string `json:"label"`
}

type QueryRequest struct{}
type QueryResponse []AlertLevelItem

type CreateRequest struct {
	Value int    `json:"value"`
	Color string `json:"color"`
	Label string `json:"label"`
}
type CreateResponse struct{}

type UpdateRequest struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
	Color string `json:"color"`
	Label string `json:"label"`
}
type UpdateResponse struct{}

type DeleteRequest struct {
	ID string `json:"id"`
}
type DeleteResponse struct{}
