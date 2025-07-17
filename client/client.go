package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"resty.dev/v3"
	"time"
)

type Response[T any] struct {
	Status int    `json:"status,omitempty"` // http 状态码
	Code   string `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   T      `json:"data,omitempty"`
}

func (resp *Response[T]) Error() string {
	if resp.Msg == "" {
		return resp.Code
	}
	return fmt.Sprintf("%s: %s", resp.Code, resp.Msg)
}

var cli *resty.Client

func SendRequest[T any](method, url string, reqData any, respData *Response[T]) error {
	resp, err := cli.R().
		SetBody(reqData).
		SetResult(respData).
		SetError(respData).
		Execute(method, url)
	if err != nil {
		return err
	}

	respData.Status = resp.StatusCode()
	if resp.IsSuccess() {
		return nil
	} else if resp.IsError() {
		return respData
	}

	return newHTTPError(resp.StatusCode(), resp.Status())
}

func GET[T any](url string, reqData any, respData *Response[T]) error {
	return SendRequest(http.MethodGet, url, reqData, respData)
}

func POST[T any](url string, reqData any, respData *Response[T]) error {
	return SendRequest(http.MethodPost, url, reqData, respData)
}

func PUT[T any](url string, reqData any, respData *Response[T]) error {
	return SendRequest(http.MethodPut, url, reqData, respData)
}

func DELETE[T any](url string, reqData any, respData *Response[T]) error {
	return SendRequest(http.MethodDelete, url, reqData, respData)
}

func Init(baseUrl, apiKey, projectID string) *resty.Client {
	cli = resty.New()
	cli.SetTimeout(20 * time.Second)
	cli.SetBaseURL(baseUrl)
	cli.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"X-API-Key":    apiKey,
		"x-ProjectID":  projectID,
	})
	cli.SetTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})
	//httptrace.ClientTrace
	return cli
}
