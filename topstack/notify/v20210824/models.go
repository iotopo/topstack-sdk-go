package v20210824

import (
	"encoding/json"
	tchttp "github.com/iotopo/topstack-sdk-go/topstack/common/http"
	"time"
)

type SendAlertRequest struct {
	*tchttp.BaseRequest
	// 接收通知的通道：sms,wechat,voicecall
	Channels []string `json:"Channels"`
	// 接收通知的手机号
	Mobile string `json:"Mobile"`
	// 设备名称
	Name string `json:"Name"`
	// 告警类型
	Type string `json:"Type"`
	// 告警等级
	Level string `json:"Level"`
	// 告警描述
	Content string `json:"Content"`
	// 告警发生时间
	ActionTime time.Time `json:"ActionTime"`
	// 告警解除时间
	RecoverTime *time.Time `json:"RecoverTime,omitempty"`
}

type SendAlertRecoverRequest SendAlertRequest

type SendAlertResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

func (r *SendAlertResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type SendAlertRecoverResponse SendAlertResponse
func (r *SendAlertRecoverResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
