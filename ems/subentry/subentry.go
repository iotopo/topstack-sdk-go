package subentry

import "topstack-sdk-go/client"

// HourlyReport 分项逐时能耗
// GET /ems/open_api/v1/report/subentry/hourly
func HourlyReport(req ReportRequest) (resp client.Response[HourlyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/subentry/hourly", req, &resp)
	return
}

// DailyReport 分项逐日能耗
// GET /ems/open_api/v1/report/subentry/daily
func DailyReport(req ReportRequest) (resp client.Response[DailyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/subentry/daily", req, &resp)
	return
}

// MonthlyReport 分项逐月能耗
// GET /ems/open_api/v1/report/subentry/monthly
func MonthlyReport(req ReportRequest) (resp client.Response[MonthlyResponse], err error) {
	err = client.GET("/ems/open_api/v1/report/subentry/monthly", req, &resp)
	return
}

// QuerySubentryList 查询指定能源分类下的分项列表
// GET /ems/open_api/v1/subentry
func QuerySubentryList(req SubentryListRequest) (resp client.Response[SubentryListResponse], err error) {
	err = client.GET("/ems/open_api/v1/subentry", req, &resp)
	return
}

// QuerySubentryDetail 能源分项详情
func QuerySubentryDetail(subentryID string) (resp client.Response[SubentryDetailResponse], err error) {
	err = client.GET("/ems/open_api/v1/subentry/"+subentryID, nil, &resp)
	return
}
