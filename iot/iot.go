package iot

import (
	"topstack-sdk-go/client"
)

func FindLast(req FindLastRequest) (resp client.Response[FindLastResponse], err error) {
	err = client.POST("/iot/open_api/v1/data/findLast", req, &resp)
	return
}

func FindLastBatch(req FindLastBatchRequest) (resp client.Response[FindLastBatchResponse], err error) {
	err = client.POST("/iot/open_api/v1/data/findLastBatch", req, &resp)
	return
}

func SetValue(req SetValueRequest) (resp client.Response[any], err error) {
	err = client.POST("/iot/open_api/v1/data/setValue", req, &resp)
	return
}

func QueryHistory(req HistoryRequest) (resp client.Response[HistoryResponse], err error) {
	err = client.POST("/iot/open_api/v1/data/query", req, &resp)
	return
}
