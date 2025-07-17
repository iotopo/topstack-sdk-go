package meter

import "topstack-sdk-go/client"

// HourlyReport 仪表逐时能耗
func HourlyReport(req ReportRequest) (resp client.Response[HourlyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/meters/hourly", req, &resp)
	return
}

// DailyReport 仪表逐日能耗
func DailyReport(req ReportRequest) (resp client.Response[DailyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/meters/daily", req, &resp)
	return
}

// MonthlyReport 仪表逐月能耗
func MonthlyReport(req ReportRequest) (resp client.Response[MonthlyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/meters/monthly", req, &resp)
	return
}

// QueryMeterList 仪表列表
func QueryMeterList(req MeterListRequest) (resp client.Response[MeterListResponse], err error) {
	err = client.GET("/ems/open_api/v1/meter", req, &resp)
	return
}

// QueryMeterDetail 仪表详情
func QueryMeterDetail(meterID string) (resp client.Response[MeterDetailResponse], err error) {
	err = client.GET("/ems/open_api/v1/meter/"+meterID, nil, &resp)
	return
}
