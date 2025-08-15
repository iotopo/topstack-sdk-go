package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HMACRoundTripper 实现 http.RoundTripper，用于 HMAC 签名
type HMACRoundTripper struct {
	Transport http.RoundTripper // 底层 Transport（默认使用 http.DefaultTransport）
	AppKey    string
	AppSecret string
}

// RoundTrip 实现 http.RoundTripper 接口
func (rt *HMACRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// 1. 读取请求 Body（后续需重新填充）
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重新填充 Body
	}

	// 2. 生成签名所需参数
	timestamp := time.Now().UTC().Format(time.RFC3339)
	nonce := generateNonce() // 生成随机 Nonce

	// 3. 计算 HMAC-SHA256 签名
	signature := computeHMAC(
		req.Method,
		req.URL.Path,
		bodyBytes,
		rt.AppKey,
		rt.AppSecret,
		timestamp,
		nonce,
	)

	// 4. 设置 Authorization 头
	authHeader := fmt.Sprintf(
		"HMAC-SHA256 %s:%s:%s:%s",
		rt.AppKey,
		signature,
		timestamp,
		nonce,
	)
	req.Header.Set("Authorization", authHeader)

	// 5. 使用底层 Transport 发送请求
	if rt.Transport == nil {
		rt.Transport = http.DefaultTransport
	}
	return rt.Transport.RoundTrip(req)
}

// computeHMAC 计算 HMAC-SHA256 签名
func computeHMAC(
	method, path string,
	body []byte,
	appKey, appSecret, timestamp, nonce string,
) string {
	// 签名内容：Method + Path + Body + AppKey + Timestamp + Nonce
	payload := fmt.Sprintf("%s%s%s%s%s%s", method, path, body, appKey, timestamp, nonce)

	// 计算 HMAC-SHA256
	h := hmac.New(sha256.New, []byte(appSecret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

// generateNonce 生成随机 Nonce（示例：UUID 或随机字符串）
func generateNonce() string {
	return "random-" + time.Now().Format("20060102150405")
}
