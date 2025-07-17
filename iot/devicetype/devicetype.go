package devicetype

import (
	"topstack-sdk-go/client"
)

func Query(req QueryRequest) (resp client.Response[QueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_type/query", req, &resp)
	return
}

func QueryPoint(req PointQueryRequest) (resp client.Response[PointQueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_type_point/query", req, &resp)
	return
} 