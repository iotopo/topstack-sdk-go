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

// Option 用于配置 resty.Client
// 例如: Debug、Logger 等
type Option func(*resty.Client)

// WithDebug 开启或关闭 Debug 模式
func WithDebug(debug bool) Option {
	return func(c *resty.Client) {
		c.SetDebug(debug)
	}
}

// WithLogger 设置自定义 Logger
func WithLogger(logger resty.Logger) Option {
	return func(c *resty.Client) {
		c.SetLogger(logger)
	}
}

// WithTimeout 设置超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *resty.Client) {
		c.SetTimeout(timeout)
	}
}

func Init(baseUrl, apiKey, projectID string, opts ...Option) *resty.Client {
	cli = resty.New()
	cli.SetTimeout(20 * time.Second)
	cli.SetBaseURL(baseUrl)
	cli.SetHeader("X-API-Key", apiKey)
	cli.SetHeader("X-ProjectID", projectID)
	cli.SetTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})
	// 应用 Option
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}

// InitClient API 密钥认证
// baseUrl 基础 URL
// appKey 应用ID
// appSecret 应用密钥
func InitClient(baseUrl, appKey, appSecret string, opts ...Option) *resty.Client {
	cli = resty.New()
	cli.SetTimeout(20 * time.Second)
	cli.SetBaseURL(baseUrl)
	cli.SetTransport(&HMACRoundTripper{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		AppKey:    appKey,
		AppSecret: appSecret,
	})
	// 应用 Option
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}
