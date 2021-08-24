package errors

import (
	"fmt"
)

type TopStackSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *TopStackSDKError) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[TopStackSDKError] Code=%s, Message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("[TopStackSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewTopStackSDKError(code, message, requestId string) error {
	return &TopStackSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *TopStackSDKError) GetCode() string {
	return e.Code
}

func (e *TopStackSDKError) GetMessage() string {
	return e.Message
}

func (e *TopStackSDKError) GetRequestId() string {
	return e.RequestId
}
