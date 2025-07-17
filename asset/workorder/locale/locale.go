package locale

import "topstack-sdk-go/client"

// List 现场工单列表
func List(req ListRequest) (resp client.Response[ListResponse], err error) {
	err = client.GET("/asset/open_api/v1/locale_work_order", req, &resp)
	return
}

// Detail 现场工单详情
func Detail(workOrderID string) (resp client.Response[DetailResponse], err error) {
	err = client.GET("/asset/open_api/v1/locale_work_order/"+workOrderID, nil, &resp)
	return
}
