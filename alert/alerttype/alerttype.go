package alerttype

import "topstack-sdk-go/client"

// List 告警类型分页查询
// GET /alert/open_api/v1/alert_type
func List(req ListRequest) (resp client.Response[ListResponse], err error) {
	err = client.GET("/alert/open_api/v1/alert_type", req, &resp)
	return
}

// Create 告警类型新增
// POST /alert/open_api/v1/alert_type
func Create(req CreateRequest) (resp client.Response[CreateResponse], err error) {
	err = client.POST("/alert/open_api/v1/alert_type", req, &resp)
	return
}

// Update 告警类型修改
// PUT /alert/open_api/v1/alert_type
func Update(req UpdateRequest) (resp client.Response[UpdateResponse], err error) {
	err = client.PUT("/alert/open_api/v1/alert_type", req, &resp)
	return
}

// Delete 告警类型删除
// DELETE /alert/open_api/v1/alert_type/{id}
func Delete(req DeleteRequest) (resp client.Response[DeleteResponse], err error) {
	err = client.DELETE("/alert/open_api/v1/alert_type/"+req.ID, req, &resp)
	return
}
