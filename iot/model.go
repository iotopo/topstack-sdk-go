package iot

import (
	"time"
)

// Data related models
type FindLastRequest struct {
	DeviceID string `json:"deviceID"`
	PointID  string `json:"pointID"`
}

type FindLastResponse struct {
	DeviceID  string      `json:"deviceID" json:"deviceID,omitempty"`
	PointID   string      `json:"pointID" json:"pointID,omitempty"`
	Value     interface{} `json:"value" json:"value,omitempty"`
	Quality   int         `json:"quality" json:"quality,omitempty"` // 0 表示正常，1 表示离线，2 表示无效
	Timestamp time.Time   `json:"timestamp" json:"timestamp"`
}

type FindLastBatchRequest []FindLastRequest

type FindLastBatchResponse []FindLastResponse

type SetValueRequest struct {
	DeviceID string `json:"deviceID"`
	PointID  string `json:"pointID"`
	Value    string `json:"value"`
}

type HistoryRequest struct {
	Points      []FindLastRequest `json:"points"`
	Start       time.Time         `json:"start" format:"date-time"`
	End         time.Time         `json:"end" format:"date-time"`
	Aggregation string            `json:"aggregation" enums:"first,last,min,max,mean"`
	Interval    string            `json:"interval" default:"5s"`
	Fill        string            `json:"fill" enums:"null,previous"`
	Offset      int32             `json:"offset" minimum:"0"`
	Limit       int32             `json:"limit" minimum:"0" maximum:"5000"`
	Order       string            `json:"order" enums:"asc,desc"` // asc(默认升序),desc
}

type HistoryResponse struct {
	Results []struct {
		DeviceID string `json:"deviceID"`
		PointID  string `json:"pointID"`
		Values   []struct {
			Value  any       `json:"value,omitempty"`
			First  any       `json:"first,omitempty"`
			Last   any       `json:"last,omitempty"`
			Max    any       `json:"max,omitempty"`
			Min    any       `json:"min,omitempty"`
			Mean   any       `json:"mean,omitempty"`
			Median any       `json:"median,omitempty"`
			Sum    any       `json:"sum,omitempty"`
			Count  any       `json:"count,omitempty"`
			Spread any       `json:"spread,omitempty"`
			Stddev any       `json:"stddev,omitempty"`
			Time   time.Time `json:"time,omitempty"`
		} `json:"values"`
	} `json:"results"`
}
