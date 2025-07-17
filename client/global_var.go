package client

// 读取变量值
// GET /open_api/v1/global_var/get_value

func QueryGlobalVarValue(req GlobalVarGetValueRequest) (resp Response[GlobalVarGetValueResponse], err error) {
	err = GET("/open_api/v1/global_var/get_value", req, &resp)
	return
}

// 更新变量值
// POST /open_api/v1/global_var/update_value

func UpdateGlobalVarValue(req GlobalVarUpdateValueRequest) (resp Response[GlobalVarUpdateValueResponse], err error) {
	err = POST("/open_api/v1/global_var/update_value", req, &resp)
	return
}
