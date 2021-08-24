package main

import (
	"fmt"
	"github.com/iotopo/topstack-sdk-go/topstack/common"
	"github.com/iotopo/topstack-sdk-go/topstack/common/errors"
	"github.com/iotopo/topstack-sdk-go/topstack/common/profile"
	notify "github.com/iotopo/topstack-sdk-go/topstack/notify/v20210824"
)

func main() {
	secretID := ""
	secretKey := ""
	credential := common.NewCredential(secretID, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "notify.iotopo.com/api/v1"
	cpf.HttpProfile.Scheme = "http"

	client, err := notify.NewClient(credential, cpf)
	if err != nil {
		panic(err)
	}

	request := notify.NewSendAlertRequest()
	resp, err := client.SendAlert(request)

	if _, ok := err.(*errors.TopStackSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		panic(err)
	}
	fmt.Printf("%s", resp.ToJsonString())
}
