package iot

import (
	"topstack-sdk-go/client"
)

func FindLast(req FindLastRequest) (resp client.ResponseData[FindLastResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func FindLastBatch(req FindLastBatchRequest) (resp client.ResponseData[FindLastBatchResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func SetValue(req SetValueRequest) (resp client.ResponseData[any], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func QueryData(req QueryDataRequest) (resp client.ResponseData[QueryDataResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func QueryDevice(req DeviceQueryRequest) (resp client.ResponseData[DeviceQueryResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func QueryDeviceProps(deviceID string) (resp client.ResponseData[DevicePropsQueryResponse], err error) {
	err = client.SendRequest(DevicePropsQueryRequest{DeviceID: deviceID}, &resp)
	return
}

func QueryDevicePoint(req DevicePointQueryRequest) (resp client.ResponseData[DevicePointQueryResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func QueryDeviceType(req DeviceTypeQueryRequest) (resp client.ResponseData[DeviceTypeQueryResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func QueryDeviceTypePoint(req DeviceTypePointQueryRequest) (resp client.ResponseData[DeviceTypePointQueryResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}

func QueryGateway(req GatewayQueryRequest) (resp client.ResponseData[GatewayQueryResponse], err error) {
	err = client.SendRequest(req, &resp)
	return
}
