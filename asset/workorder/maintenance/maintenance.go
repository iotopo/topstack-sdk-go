package maintenance

import "topstack-sdk-go/client"

// List 维护工单列表
func List(req ListRequest) (resp client.Response[ListResponse], err error) {
	err = client.GET("/asset/open_api/v1/maintenance_work_order", req, &resp)
	return
}

// Detail 维护工单详情
func Detail(workOrderID string) (resp client.Response[DetailResponse], err error) {
	err = client.GET("/asset/open_api/v1/maintenance_work_order/"+workOrderID, nil, &resp)
	return
}
