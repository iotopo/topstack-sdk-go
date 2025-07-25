package devicetype

import (
	"topstack-sdk-go/client"
)

func Query(req QueryRequest) (resp client.Response[QueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_type/query", req, &resp)
	return
}

func Create(req CreateRequest) (resp client.Response[CreateResponse], err error) {
	err = client.POST("/iot/open_api/v1/device_type", req, &resp)
	return
}

func Modify(req ModifyRequest) (resp client.Response[any], err error) {
	err = client.PUT("/iot/open_api/v1/device_type", req, &resp)
	return
}

func Delete(id string) (resp client.Response[any], err error) {
	url := "/iot/open_api/v1/device_type/" + id
	err = client.DELETE(url, nil, &resp)
	return
}

func QueryPoint(req PointQueryRequest) (resp client.Response[PointQueryResponse], err error) {
	err = client.GET("/iot/open_api/v1/device_type_point/query", req, &resp)
	return
}

func CreatePoint(req CreatePointRequest) (resp client.Response[any], err error) {
	err = client.POST("/iot/open_api/v1/device_type_point", req, &resp)
	return
}

func UpdatePoint(req UpdatePointRequest) (resp client.Response[any], err error) {
	err = client.PUT("/iot/open_api/v1/device_type_point", req, &resp)
	return
}

func DeletePoint(deviceTypeID, pointID string) (resp client.Response[any], err error) {
	url := "/iot/open_api/v1/device_type_point/" + deviceTypeID + "/" + pointID
	err = client.DELETE(url, nil, &resp)
	return
}

func DeleteAllPoints(deviceTypeID string) (resp client.Response[any], err error) {
	url := "/iot/open_api/v1/device_type_point/" + deviceTypeID
	err = client.DELETE(url, nil, &resp)
	return
}

func BatchCreate(req BatchCreateRequest) (resp client.Response[any], err error) {
	err = client.POST("/iot/open_api/v1/device_type/batch", req, &resp)
	return
}
