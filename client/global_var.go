package client

import "net/http"

// 读取变量值
// GET /open_api/v1/global_var/get_value

type GlobalVarGetValueRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func (r GlobalVarGetValueRequest) Method() string { return http.MethodGet }
func (r GlobalVarGetValueRequest) Url() string    { return "/open_api/v1/global_var/get_value" }

type GlobalVarValue struct {
	Namespace string      `json:"namespace"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value"`
	Type      string      `json:"type"`
	Time      string      `json:"time"`
}

type GlobalVarGetValueResponse struct {
	Success bool           `json:"success"`
	Data    GlobalVarValue `json:"data"`
}

func QueryGlobalVarValue(req GlobalVarGetValueRequest) (resp GlobalVarGetValueResponse, err error) {
	err = SendRequest(req, &resp)
	return
}

// 更新变量值
// POST /open_api/v1/global_var/update_value

type GlobalVarUpdateValueRequest struct {
	Namespace string      `json:"namespace"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value"`
}

func (r GlobalVarUpdateValueRequest) Method() string { return http.MethodPost }
func (r GlobalVarUpdateValueRequest) Url() string    { return "/open_api/v1/global_var/update_value" }

type GlobalVarUpdateValueResponse struct {
	Success bool `json:"success"`
}

func UpdateGlobalVarValue(req GlobalVarUpdateValueRequest) (resp GlobalVarUpdateValueResponse, err error) {
	err = SendRequest(req, &resp)
	return
} 