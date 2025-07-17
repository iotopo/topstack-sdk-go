package client

// 读取变量值
// GET /open_api/v1/global_var/get_value

type GlobalVarGetValueRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type GlobalVarValue struct {
	Namespace string      `json:"namespace"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value"`
	Type      string      `json:"type"`
	Time      string      `json:"time"`
}

type GlobalVarGetValueResponse GlobalVarValue

// 更新变量值
// POST /open_api/v1/global_var/update_value

type GlobalVarUpdateValueRequest struct {
	Namespace string      `json:"namespace"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value"`
}

type GlobalVarUpdateValueResponse struct{} 