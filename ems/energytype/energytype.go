package energytype

import "topstack-sdk-go/client"

// QueryEnergyTypeList 能源类型列表
func QueryEnergyTypeList(req EnergyTypeListRequest) (resp client.Response[EnergyTypeListResponse], err error) {
	err = client.GET("/ems/open_api/v1/energy_type", req, &resp)
	return
}

// QueryEnergyTypeDetail 能源类型详情
func QueryEnergyTypeDetail(energyTypeID string) (resp client.Response[EnergyTypeDetailResponse], err error) {
	err = client.GET("/ems/open_api/v1/energy_type/"+energyTypeID, nil, &resp)
	return
}
