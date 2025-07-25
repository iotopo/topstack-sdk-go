package device

import (
	"topstack-sdk-go/client"
)

func Query(req QueryRequest) (resp client.Response[QueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device/query", req, &resp)
	return
}

func QueryProps(deviceID string) (resp client.Response[PropsQueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device/"+deviceID+"/props", nil, &resp)
	return
}

func Create(req CreateRequest) (resp client.Response[CreateResponse], err error) {
	err = client.POST("/iot/open_api/v1/device", req, &resp)
	return
}

func Modify(req ModifyRequest) (resp client.Response[any], err error) {
	err = client.PUT("/iot/open_api/v1/device", req, &resp)
	return
}

func Delete(id string) (resp client.Response[any], err error) {
	url := "/iot/open_api/v1/device/" + id
	err = client.DELETE(url, nil, &resp)
	return
}

func QueryPoint(req PointQueryRequest) (resp client.Response[PointQueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_point/query", req, &resp)
	return
}

func BatchCreate(req BatchCreateRequest) (resp client.Response[any], err error) {
	err = client.POST("/iot/open_api/v1/device/batch", req, &resp)
	return
}
