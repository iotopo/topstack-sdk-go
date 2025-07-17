package sector

import "time"

type SectorListRequest struct{}

type SectorListItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	ParentID    string `json:"parentID,omitempty"`
	Description string `json:"description,omitempty"`
}

type SectorListResponse []SectorListItem

type SectorDetailResponse SectorListItem

type ReportRequest struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	EnergyTypeID string    `json:"energyTypeID"`
	SectorID     string    `json:"sectorID,omitempty"`
	SectorCode   string    `json:"sectorCode,omitempty"`
	Round        int       `json:"round,omitempty"`
}

type Value struct {
	Time  string  `json:"time"`
	Value *string `json:"value"`
}

type ReportItem struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Code   string  `json:"code"`
	Total  *string `json:"total"`
	Values []Value `json:"values"`
}

type HourlyResponse []ReportItem

type DailyResponse = HourlyResponse

type MonthlyResponse = HourlyResponse
