package sector

import "topstack-sdk-go/client"

// HourlyReport 用能单元逐时能耗
func HourlyReport(req ReportRequest) (resp client.Response[HourlyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/sector/hourly", req, &resp)
	return
}

// DailyReport 用能单元逐日能耗
func DailyReport(req ReportRequest) (resp client.Response[DailyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/sector/daily", req, &resp)
	return
}

// MonthlyReport 用能单元逐月能耗
func MonthlyReport(req ReportRequest) (resp client.Response[MonthlyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/sector/monthly", req, &resp)
	return
}

// QuerySectorList 用能单元列表
func QuerySectorList(req SectorListRequest) (resp client.Response[SectorListResponse], err error) {
	err = client.GET("/ems/open_api/v1/sector", req, &resp)
	return
}

// QuerySectorDetail 用能单元详情
func QuerySectorDetail(sectorID string) (resp client.Response[SectorDetailResponse], err error) {
	err = client.GET("/ems/open_api/v1/sector/"+sectorID, nil, &resp)
	return
}
