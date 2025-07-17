package subentry

import "time"

type SubentryListRequest struct {
	EnergyTypeID string `json:"energyTypeID,omitempty"`
}

type SubentryListItem struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	EnergyTypeID string `json:"energyTypeID"`
	ParentID     string `json:"parentID,omitempty"`
}

type SubentryListResponse []SubentryListItem

type SubentryDetailResponse SubentryListItem

type ReportRequest struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	SectorID     string    `json:"sectorID"`
	EnergyTypeID string    `json:"energyTypeID"`
	SubentryID   string    `json:"subentryID,omitempty"`
	SubentryCode string    `json:"subentryCode,omitempty"`
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
