package datav

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"time"
)

type Param struct {
	Name  string
	Value string
}

// GetPageUrl 生成组态画面访问地址
func GetPageUrl(baseUrl, pageID, token, username, op string, params ...Param) (string, error) {
	ts := time.Now().UnixMilli()

	h := hmac.New(sha256.New, []byte(token))
	io.WriteString(h, pageID)
	io.WriteString(h, "|")
	io.WriteString(h, fmt.Sprintf("%d", ts))
	io.WriteString(h, "|")
	io.WriteString(h, username)
	io.WriteString(h, "|")
	io.WriteString(h, op)

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	u, err := url.Parse(fmt.Sprintf("%s/topv/public/viewer/%s?", baseUrl, pageID))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("time", fmt.Sprintf("%d", ts))
	q.Set("username", username)
	q.Set("op", op)
	q.Set("signature", signature)
	for _, param := range params {
		q.Set(param.Name, param.Value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// GetGraphEditorUrl 生成组态编辑器访问地址
func GetGraphEditorUrl(baseUrl, projectID, token string, params ...Param) (string, error) {
	ts := time.Now().UnixMilli()

	h := hmac.New(sha256.New, []byte(token))
	io.WriteString(h, fmt.Sprintf("%d", ts))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	u, err := url.Parse(fmt.Sprintf("%s/topv/public/editor/%s?", baseUrl, projectID))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("time", fmt.Sprintf("%d", ts))
	q.Set("signature", signature)
	for _, param := range params {
		q.Set(param.Name, param.Value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// GetComposeEditorUrl 生成门户编辑器访问地址
func GetComposeEditorUrl(baseUrl, projectID, token string, params ...Param) (string, error) {
	ts := time.Now().UnixMilli()

	h := hmac.New(sha256.New, []byte(token))
	io.WriteString(h, fmt.Sprintf("%d", ts))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	u, err := url.Parse(fmt.Sprintf("%s/topv/public/compose/%s?", baseUrl, projectID))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("time", fmt.Sprintf("%d", ts))
	q.Set("signature", signature)
	for _, param := range params {
		q.Set(param.Name, param.Value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// GetPortalUrl 生成 PC 端数据门户访问地址
func GetPortalUrl(baseUrl, projectID, token string, params ...Param) (string, error) {
	ts := time.Now().UnixMilli()

	h := hmac.New(sha256.New, []byte(token))
	io.WriteString(h, fmt.Sprintf("%d", ts))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	u, err := url.Parse(fmt.Sprintf("%s/topv/public/portal/%s?", baseUrl, projectID))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("time", fmt.Sprintf("%d", ts))
	q.Set("signature", signature)
	for _, param := range params {
		q.Set(param.Name, param.Value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// GetMobileUrl 生成移动端访问地址
func GetMobileUrl(baseUrl, projectID, token string, params ...Param) (string, error) {
	ts := time.Now().UnixMilli()

	h := hmac.New(sha256.New, []byte(token))
	io.WriteString(h, fmt.Sprintf("%d", ts))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	u, err := url.Parse(fmt.Sprintf("%s/topv/public/mobile/%s?", baseUrl, projectID))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("time", fmt.Sprintf("%d", ts))
	q.Set("signature", signature)
	for _, param := range params {
		q.Set(param.Name, param.Value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
