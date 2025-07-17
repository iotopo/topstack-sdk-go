package gateway

import (
	"topstack-sdk-go/client"
)

func Query(req QueryRequest) (resp client.Response[QueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/gateway/query", req, &resp)
	return
} 