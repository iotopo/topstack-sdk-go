package v20210824

import (
	"github.com/iotopo/topstack-sdk-go/topstack/common"
	tchttp "github.com/iotopo/topstack-sdk-go/topstack/common/http"
	"github.com/iotopo/topstack-sdk-go/topstack/common/profile"
)

const APIVersion = "2021-08-24"

type Client struct {
	common.Client
}

func (c *Client) SendAlert(request *SendAlertRequest) (response *SendAlertResponse, err error) {
	if request == nil {
		request = NewSendAlertRequest()
	}
	response = NewSendAlertResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) SendAlertRecover(request *SendAlertRecoverRequest) (response *SendAlertRecoverResponse, err error) {
	if request == nil {
		request = NewSendAlertRecoverRequest()
	}
	response = NewSendAlertRecoverResponse()
	err = c.Send(request, response)
	return
}

func NewClient(credential common.CredentialIface, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential).WithProfile(clientProfile).Init()
	return
}

func NewSendAlertRequest() (request *SendAlertRequest) {
	request = &SendAlertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("notify", APIVersion, "SendAlert")
	request.SetHttpMethod("POST")
	return
}

func NewSendAlertResponse() (response *SendAlertResponse) {
	response = &SendAlertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func NewSendAlertRecoverRequest() (request *SendAlertRecoverRequest) {
	request = &SendAlertRecoverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("notify", APIVersion, "SendAlertRecover")
	return
}

func NewSendAlertRecoverResponse() (response *SendAlertRecoverResponse) {
	response = &SendAlertRecoverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}
