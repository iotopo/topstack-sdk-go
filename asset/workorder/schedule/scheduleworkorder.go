package schedule

import "topstack-sdk-go/client"

// List 计划工单列表
func List(req ListRequest) (resp client.Response[ListResponse], err error) {
	err = client.GET("/asset/open_api/v1/schedule_work_order", req, &resp)
	return
}

// Detail 计划工单详情
func Detail(req DetailRequest) (resp client.Response[DetailResponse], err error) {
	err = client.GET("/asset/open_api/v1/schedule_work_order/"+req.ID, req, &resp)
	return
}
