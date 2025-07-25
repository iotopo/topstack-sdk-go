package device_type_group

import (
	"topstack-sdk-go/client"
)

func Query(req QueryRequest) (resp client.Response[QueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_type_group/all", req, &resp)
	return
}

func Create(req CreateRequest) (resp client.Response[CreateResponse], err error) {
	err = client.POST("/iot/open_api/v1/device_type_group", req, &resp)
	return
}

func Modify(req ModifyRequest) (resp client.Response[any], err error) {
	err = client.PUT("/iot/open_api/v1/device_type_group", req, &resp)
	return
}

func Delete(req DeleteRequest) (resp client.Response[any], err error) {
	err = client.POST("/iot/open_api/v1/device_type_group/delete", req, &resp)
	return
}

func BatchCreate(req BatchCreateRequest) (resp client.Response[any], err error) {
	err = client.POST("/iot/open_api/v1/device_type_group/batch", req, &resp)
	return
} 