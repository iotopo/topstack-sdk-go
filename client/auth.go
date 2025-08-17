package client

import (
	"fmt"
	"time"
)

//var tokenMutex sync.Mutex
//var accessToken string
//var expiresAt time.Time

func (cli *Client) getAccessToken() error {
	cli.tokenMutex.Lock()
	defer cli.tokenMutex.Unlock()

	// 检查令牌是否还有效
	if cli.AccessToken != "" && time.Now().Before(cli.ExpiresAt) {
		return nil
	}

	var tokenResp struct {
		Code              string `json:"code"`
		Msg               string `json:"msg"`
		TenantAccessToken string `json:"access_token"`
		Expire            int    `json:"expire"`
	}

	res, err := cli.cli.R().
		SetForceResponseContentType("application/json").
		SetBody(map[string]string{
			"app_id":     cli.AppID,
			"app_secret": cli.AppSecret,
		}).
		SetResult(&tokenResp).
		Post("/open_api/v1/auth/access_token")

	if err != nil {
		return fmt.Errorf("获取访问令牌失败: %w", err)
	}

	if res.IsSuccess() {
		if tokenResp.Code != "" {
			return fmt.Errorf("获取访问令牌失败: %s, %s", tokenResp.Code, tokenResp.Msg)
		}
		cli.AccessToken = tokenResp.TenantAccessToken
		cli.ExpiresAt = time.Now().Add(time.Duration(tokenResp.Expire-300) * time.Second) // 提前5分钟过期
		return nil
	} else {
		return fmt.Errorf("获取访问令牌失败: %s", res.Status())
	}
}
