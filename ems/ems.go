package ems

import (
	"net/http"
	"topstack-sdk-go/client"
)

// 仪表列表
// GET /ems/open_api/v1/meter

type MeterListRequest struct {
	EnergyTypeID string `json:"energyTypeID,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	Tags         string `json:"tags,omitempty"`
}

func (MeterListRequest) Method() string { return http.MethodGet }
func (MeterListRequest) Url() string    { return "/ems/open_api/v1/meter" }

type MeterListItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	EnergyTypeID string `json:"energyTypeID"`
	Description string `json:"description,omitempty"`
}

type MeterListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Total int             `json:"total"`
		Items []MeterListItem `json:"items"`
	} `json:"data"`
}

func QueryMeterList(req MeterListRequest) (resp MeterListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 能源类型列表
// GET /ems/open_api/v1/energy_type

type EnergyTypeListRequest struct{}

func (EnergyTypeListRequest) Method() string { return http.MethodGet }
func (EnergyTypeListRequest) Url() string    { return "/ems/open_api/v1/energy_type" }

type EnergyTypeListItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type EnergyTypeListResponse struct {
	Success bool                `json:"success"`
	Data    []EnergyTypeListItem `json:"data"`
}

func QueryEnergyTypeList(req EnergyTypeListRequest) (resp EnergyTypeListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 用能单元列表
// GET /ems/open_api/v1/sector

type SectorListRequest struct{}

func (SectorListRequest) Method() string { return http.MethodGet }
func (SectorListRequest) Url() string    { return "/ems/open_api/v1/sector" }

type SectorListItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	ParentID    string `json:"parentID,omitempty"`
	Description string `json:"description,omitempty"`
}

type SectorListResponse struct {
	Success bool            `json:"success"`
	Data    []SectorListItem `json:"data"`
}

func QuerySectorList(req SectorListRequest) (resp SectorListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 能源分项列表
// GET /ems/open_api/v1/subentry

type SubentryListRequest struct {
	EnergyTypeID string `json:"energyTypeID,omitempty"`
}

func (SubentryListRequest) Method() string { return http.MethodGet }
func (SubentryListRequest) Url() string    { return "/ems/open_api/v1/subentry" }

type SubentryListItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	EnergyTypeID string `json:"energyTypeID"`
	ParentID    string `json:"parentID,omitempty"`
}

type SubentryListResponse struct {
	Success bool              `json:"success"`
	Data    []SubentryListItem `json:"data"`
}

func QuerySubentryList(req SubentryListRequest) (resp SubentryListResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 仪表详情
// GET /ems/open_api/v1/meter/:id

type MeterDetailRequest struct {
	ID string
}

func (r MeterDetailRequest) Method() string { return http.MethodGet }
func (r MeterDetailRequest) Url() string    { return "/ems/open_api/v1/meter/" + r.ID }

type MeterDetailResponse struct {
	Success bool        `json:"success"`
	Data    MeterListItem `json:"data"`
}

func QueryMeterDetail(req MeterDetailRequest) (resp MeterDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 能源类型详情
// GET /ems/open_api/v1/energy_type/:id

type EnergyTypeDetailRequest struct {
	ID string
}

func (r EnergyTypeDetailRequest) Method() string { return http.MethodGet }
func (r EnergyTypeDetailRequest) Url() string    { return "/ems/open_api/v1/energy_type/" + r.ID }

type EnergyTypeDetailResponse struct {
	Success bool              `json:"success"`
	Data    EnergyTypeListItem `json:"data"`
}

func QueryEnergyTypeDetail(req EnergyTypeDetailRequest) (resp EnergyTypeDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 用能单元详情
// GET /ems/open_api/v1/sector/:id

type SectorDetailRequest struct {
	ID string
}

func (r SectorDetailRequest) Method() string { return http.MethodGet }
func (r SectorDetailRequest) Url() string    { return "/ems/open_api/v1/sector/" + r.ID }

type SectorDetailResponse struct {
	Success bool           `json:"success"`
	Data    SectorListItem `json:"data"`
}

func QuerySectorDetail(req SectorDetailRequest) (resp SectorDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
}

// 能源分项详情
// GET /ems/open_api/v1/subentry/:id

type SubentryDetailRequest struct {
	ID string
}

func (r SubentryDetailRequest) Method() string { return http.MethodGet }
func (r SubentryDetailRequest) Url() string    { return "/ems/open_api/v1/subentry/" + r.ID }

type SubentryDetailResponse struct {
	Success bool             `json:"success"`
	Data    SubentryListItem `json:"data"`
}

func QuerySubentryDetail(req SubentryDetailRequest) (resp SubentryDetailResponse, err error) {
	err = client.SendRequest(req, &resp)
	return
} 