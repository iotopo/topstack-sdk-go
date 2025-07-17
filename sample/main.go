package main

import (
	"fmt"
	"time"
	"topstack-sdk-go/client"
	"topstack-sdk-go/datav"
	"topstack-sdk-go/iot"
)

func main() {
	apiKey := "x248bz0728u0hlhn"
	projectID := "iotopo"
	cli := client.Init("http://localhost:8000", apiKey, projectID)
	cli.SetDebug(true)

	// 查询单测点实时值
	{
		_, err := iot.FindLast(iot.FindLastRequest{DeviceID: "dev1", PointID: "v1"})
		if err != nil {
			panic(err)
		}
	}

	// 批量查询多测点实时值
	{
		_, err := iot.FindLastBatch(iot.FindLastBatchRequest{{DeviceID: "dev1", PointID: "v1"}})
		if err != nil {
			panic(err)
		}
	}

	// 控制指令下发
	{
		_, err := iot.SetValue(iot.SetValueRequest{DeviceID: "dev1", PointID: "V4", Value: "2"})
		if err != nil {
			panic(err)
		}
	}

	{
		_, err := iot.QueryData(iot.QueryDataRequest{
			Points:      []iot.FindLastRequest{{DeviceID: "dev1", PointID: "v1"}},
			Start:       time.Now().Add(-10 * time.Minute),
			End:         time.Now(),
			Interval:    "10s",
			Aggregation: "last",
		})
		if err != nil {
			panic(err)
		}
	}

	{
		_, err := iot.QueryDevice(iot.DeviceQueryRequest{Search: "dev1", PageNum: 1, PageSize: 10})
		if err != nil {
			panic(err)
		}
	}

	{
		_, err := iot.QueryDeviceProps("dev1")
		if err != nil {
			panic(err)
		}
	}

	{
		pageUrl, err := datav.GetPageUrl("http://localhost:8000", "ct9avosgj6lmas2l2hl0", "8061337fe3e34be1bf0ee1a4e15d3a63", "", "")
		if err != nil {
			panic(err)
		}
		fmt.Println(pageUrl)
	}
}
