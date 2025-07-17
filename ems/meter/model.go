package meter

import "time"

type MeterListRequest struct {
	EnergyTypeID string `json:"energyTypeID,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	Tags         string `json:"tags,omitempty"`
}

type MeterListItem struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	EnergyTypeID string `json:"energyTypeID"`
	Description  string `json:"description,omitempty"`
}

type MeterListResponse struct {
	Total int             `json:"total"`
	Items []MeterListItem `json:"items"`
}

type MeterDetailResponse MeterListItem

type ReportRequest struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	MeterID      string    `json:"meterID,omitempty"`
	MeterCode    string    `json:"meterCode,omitempty"`
	EnergyTypeID string    `json:"energyTypeID,omitempty"`
	PageSize     int       `json:"pageSize,omitempty"`
	PageNum      int       `json:"pageNum,omitempty"`
	Tags         string    `json:"tags,omitempty"`
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

type HourlyResponse struct {
	Total int          `json:"total"`
	Items []ReportItem `json:"items"`
}

type DailyResponse ReportRequest

type MonthlyResponse ReportRequest
