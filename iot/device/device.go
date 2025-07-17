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

func QueryPoint(req PointQueryRequest) (resp client.Response[PointQueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_point/query", req, &resp)
	return
}
