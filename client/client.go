package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"resty.dev/v3"
	"time"
)

type RequestData interface {
	Method() string
	Url() string
}

type ResponseData[T any] struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
}

func (resp *ResponseData[T]) Error() string {
	return fmt.Sprintf("%s: %s", resp.Code, resp.Msg)
}

type ErrorResponse struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (resp *ErrorResponse) Error() string {
	if resp.Msg == "" {
		return resp.Code
	}
	return fmt.Sprintf("%s: %s", resp.Code, resp.Msg)
}

var cli *resty.Client

//type Client struct {
//	apiKey    string
//	projectID string
//	baseUrl   string
//	//client    *http.Client
//	client *resty.Client
//}

//	func (c *Client) SendRequest(data RequestData, respData any) error {
//		var body io.Reader
//		if data != nil {
//			b, err := json.Marshal(data)
//			if err != nil {
//				return err
//			}
//			body = bytes.NewReader(b)
//		}
//
//		req, err := http.NewRequest(data.Method(), c.baseUrl+data.Url(), body)
//		if err != nil {
//			return err
//		}
//		req.Header.Set("X-API-Key", c.apiKey)
//		req.Header.Set("X-ProjectID", c.projectID)
//		req.Header.Set("Content-Type", "application/json")
//		resp, err := c.client.Do(req)
//		if err != nil {
//			return err
//		}
//		defer resp.Body.Close()
//		//return io.ReadAll(resp.Body)
//		b, err := io.ReadAll(resp.Body)
//		if err != nil {
//			return err
//		}
//		if resp.StatusCode != http.StatusOK {
//			log.Println(string(b))
//			return newHTTPError(resp.StatusCode, resp.Status)
//		}
//		if respData != nil {
//			err = json.Unmarshal(b, respData)
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	}
func SendRequest(reqData RequestData, respData any) error {
	resp, err := cli.R().
		SetBody(reqData).
		SetResult(respData).
		SetError(respData).
		Execute(reqData.Method(), reqData.Url())
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		errResp := resp.Error().(*ErrorResponse)
		if errResp.Code != "" {
			return errResp
		}
		return newHTTPError(resp.StatusCode(), resp.Status())
	}
	return nil
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
