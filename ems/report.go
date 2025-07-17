package ems

import (
	"net/http"
	"topstack-sdk-go/client"
)

type MetersReportRequest struct {
	Start        string `json:"start"`
	End          string `json:"end"`
	MeterID      string `json:"meterID,omitempty"`
	MeterCode    string `json:"meterCode,omitempty"`
	EnergyTypeID string `json:"energyTypeID,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	Tags         string `json:"tags,omitempty"`
}

// 仪表逐时能耗
// GET /ems/open_api/v1/report/meters/hourly
// 只实现部分主要字段，详细字段可根据需要补充

type MetersHourlyReportRequest MetersReportRequest

func (MetersHourlyReportRequest) Method() string { return http.MethodGet }
func (MetersHourlyReportRequest) Url() string    { return "/ems/open_api/v1/report/meters/hourly" }

type MeterValue struct {
	Time  string  `json:"time"`
	Value *string `json:"value"`
}

type MeterReportItem struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Code   string       `json:"code"`
	Total  *string      `json:"total"`
	Values []MeterValue `json:"values"`
}

type MetersHourlyReportResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int               `json:"total"`
		Items []MeterReportItem `json:"items"`
	} `json:"data"`
}

func QueryMetersHourlyReport(req MetersHourlyReportRequest) (resp MetersHourlyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 仪表逐日能耗
// GET /ems/open_api/v1/report/meters/daily

type MetersDailyReportRequest MetersReportRequest

func (MetersDailyReportRequest) Method() string { return http.MethodGet }
func (MetersDailyReportRequest) Url() string    { return "/ems/open_api/v1/report/meters/daily" }

type MetersDailyReportResponse MetersReportRequest

func QueryMetersDailyReport(req MetersDailyReportRequest) (resp MetersDailyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 仪表逐月能耗
// GET /ems/open_api/v1/report/meters/monthly

type MetersMonthlyReportRequest = MetersReportRequest

func (MetersMonthlyReportRequest) Method() string { return http.MethodGet }
func (MetersMonthlyReportRequest) Url() string    { return "/ems/open_api/v1/report/meters/monthly" }

type MetersMonthlyReportResponse MetersReportRequest

func QueryMetersMonthlyReport(req MetersMonthlyReportRequest) (resp MetersMonthlyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

type SectorReportRequest struct {
	Start        string `json:"start"`
	End          string `json:"end"`
	EnergyTypeID string `json:"energyTypeID"`
	SectorID     string `json:"sectorID,omitempty"`
	SectorCode   string `json:"sectorCode,omitempty"`
	Round        int    `json:"round,omitempty"`
}

// 用能单元逐时能耗
// GET /ems/open_api/v1/report/sector/hourly

type SectorHourlyReportRequest SectorReportRequest

func (SectorHourlyReportRequest) Method() string { return http.MethodGet }
func (SectorHourlyReportRequest) Url() string    { return "/ems/open_api/v1/report/sector/hourly" }

type SectorReportItem struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Code   string       `json:"code"`
	Total  *string      `json:"total"`
	Values []MeterValue `json:"values"`
}
type SectorHourlyReportResponse struct {
	Success bool               `json:"success"`
	Data    []SectorReportItem `json:"data"`
}

func QuerySectorHourlyReport(req SectorHourlyReportRequest) (resp SectorHourlyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 用能单元逐日能耗
// GET /ems/open_api/v1/report/sector/daily

type SectorDailyReportRequest SectorReportRequest

func (SectorDailyReportRequest) Method() string { return http.MethodGet }
func (SectorDailyReportRequest) Url() string    { return "/ems/open_api/v1/report/sector/daily" }

type SectorDailyReportResponse = SectorHourlyReportResponse

func QuerySectorDailyReport(req SectorDailyReportRequest) (resp SectorDailyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 用能单元逐月能耗
// GET /ems/open_api/v1/report/sector/monthly

type SectorMonthlyReportRequest SectorReportRequest

func (SectorMonthlyReportRequest) Method() string { return http.MethodGet }
func (SectorMonthlyReportRequest) Url() string { return "/ems/open_api/v1/report/sector/monthly" }

type SectorMonthlyReportResponse = SectorHourlyReportResponse

func QuerySectorMonthlyReport(req SectorMonthlyReportRequest) (resp SectorMonthlyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

type SubentryReportRequest struct {
	Start        string `json:"start"`
	End          string `json:"end"`
	SectorID     string `json:"sectorID"`
	EnergyTypeID string `json:"energyTypeID"`
	SubentryID   string `json:"subentryID,omitempty"`
	SubentryCode string `json:"subentryCode,omitempty"`
	Round        int    `json:"round,omitempty"`
}

// 分项逐时能耗
// GET /ems/open_api/v1/report/subentry/hourly

type SubentryHourlyReportRequest SubentryReportRequest

func (SubentryHourlyReportRequest) Method() string { return http.MethodGet }
func (SubentryHourlyReportRequest) Url() string    { return "/ems/open_api/v1/report/subentry/hourly" }

type SubentryReportItem = SectorReportItem
type SubentryHourlyReportResponse struct {
	Success bool                 `json:"success"`
	Data    []SubentryReportItem `json:"data"`
}

func QuerySubentryHourlyReport(req SubentryHourlyReportRequest) (resp SubentryHourlyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 分项逐日能耗
// GET /ems/open_api/v1/report/subentry/daily

type SubentryDailyReportRequest SubentryReportRequest

func (SubentryDailyReportRequest) Method() string { return http.MethodGet }
func (SubentryDailyReportRequest) Url() string    { return "/ems/open_api/v1/report/subentry/daily" }

type SubentryDailyReportResponse = SubentryHourlyReportResponse

func QuerySubentryDailyReport(req SubentryDailyReportRequest) (resp SubentryDailyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 分项逐月能耗
// GET /ems/open_api/v1/report/subentry/monthly

type SubentryMonthlyReportRequest SubentryReportRequest

func (SubentryMonthlyReportRequest) Method() string { return http.MethodGet }
func (SubentryMonthlyReportRequest) Url() string    { return "/ems/open_api/v1/report/subentry/monthly" }

type SubentryMonthlyReportResponse = SubentryHourlyReportResponse

func QuerySubentryMonthlyReport(req SubentryMonthlyReportRequest) (resp SubentryMonthlyReportResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}
