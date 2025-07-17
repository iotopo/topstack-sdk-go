package alertlevel

import "topstack-sdk-go/client"

// Query 告警等级查询
// GET /alert/open_api/v1/alert_level
func Query() (resp client.Response[QueryResponse], err error) {
	err = client.GET("/alert/open_api/v1/alert_level", QueryRequest{}, &resp)
	return
}

// Create 告警等级新增
// POST /alert/open_api/v1/alert_level
func Create(req CreateRequest) (resp client.Response[CreateResponse], err error) {
	err = client.POST("/alert/open_api/v1/alert_level", req, &resp)
	return
}

// Update 告警等级修改
// PUT /alert/open_api/v1/alert_level
func Update(req UpdateRequest) (resp client.Response[UpdateResponse], err error) {
	err = client.PUT("/alert/open_api/v1/alert_level", req, &resp)
	return
}

// Delete 告警等级删除
// DELETE /alert/open_api/v1/alert_level
func Delete(req DeleteRequest) (resp client.Response[DeleteResponse], err error) {
	err = client.DELETE("/alert/open_api/v1/alert_level", req, &resp)
	return
}
