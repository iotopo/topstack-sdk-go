package energytype

type EnergyTypeListRequest struct{}

type EnergyTypeListItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type EnergyTypeListResponse []EnergyTypeListItem

type EnergyTypeDetailResponse EnergyTypeListItem
