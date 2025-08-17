package main

import (
	"fmt"
	"time"
	"topstack-sdk-go/alert/alertlevel"
	"topstack-sdk-go/client"
	"topstack-sdk-go/datav"
	"topstack-sdk-go/iot"
	"topstack-sdk-go/iot/device"
	"topstack-sdk-go/iot/device_type_group"
	"topstack-sdk-go/iot/devicetype"
)

func main() {
	//apiKey := "8mphozy98fkor6iu"
	//projectID := "iotopo"
	//client.Init("http://localhost:8000", apiKey, projectID, client.WithDebug(true))

	//appID := "ba4c3c7be64a44fba37c1cfaae0abb61"
	//appSecret := "ba4c3c7be64a44fba37c1cfaae0abb61"
	appID := "d55864f766c34d0383243eaacd8e01f7"
	appSecret := "34bc004307e8458fae7beb1d9197e759"
	client.InitClient("http://localhost:8000", appID, appSecret, client.WithDebug(true))

	{
		resp, err := alertlevel.Query()
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}

	// 查询单测点实时值
	{
		resp, err := iot.FindLast(iot.FindLastRequest{DeviceID: "Dev1", PointID: "V1"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s.%s=%v(%v)", resp.Data.DeviceID, resp.Data.PointID, resp.Data.Value, resp.Data.Timestamp)
	}

	// 批量查询多测点实时值
	{
		_, err := iot.FindLastBatch(iot.FindLastBatchRequest{{DeviceID: "Dev1", PointID: "V1"}})
		if err != nil {
			panic(err)
		}
	}

	// 控制指令下发
	{
		_, err := iot.SetValue(iot.SetValueRequest{DeviceID: "Dev1", PointID: "V4", Value: "2"})
		if err != nil {
			panic(err)
		}
	}

	{
		_, err := iot.QueryHistory(iot.HistoryRequest{
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
		_, err := device.Query(device.QueryRequest{Search: "dev1", PageNum: 1, PageSize: 10})
		if err != nil {
			panic(err)
		}
	}

	{
		_, err := device.QueryProps("dev1")
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

	// device_type_group 示例
	{
		resp, err := device_type_group.Query(device_type_group.QueryRequest{PageNum: 1, PageSize: 10})
		if err != nil {
			panic(err)
		}
		fmt.Println("device_type_group.Query:", resp)
	}

	{
		resp, err := device_type_group.Create(device_type_group.CreateRequest{
			Name:        "测试分组",
			Code:        "test-group",
			Description: "SDK创建测试",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("device_type_group.Create:", resp)
	}

	{
		resp, err := device_type_group.Modify(device_type_group.ModifyRequest{
			ID:          "group-id",
			Name:        "新分组名",
			Code:        "new-code",
			Description: "修改描述",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("device_type_group.Modify:", resp)
	}

	{
		resp, err := device_type_group.Delete(device_type_group.DeleteRequest{ID: "group-id"})
		if err != nil {
			panic(err)
		}
		fmt.Println("device_type_group.Delete:", resp)
	}

	// device_type 示例
	{
		resp, err := devicetype.Create(devicetype.CreateRequest{
			Name:        "测试类型",
			Description: "SDK创建测试类型",
			Asset:       true,
			Perf:        false,
			Code:        "test-type",
			Icon:        "icon-test",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.Create:", resp)
	}

	{
		resp, err := devicetype.Modify(devicetype.ModifyRequest{
			ID:          "type-id",
			Name:        "新类型名",
			Description: "修改描述",
			Asset:       false,
			Perf:        true,
			Code:        nil,
			Icon:        "icon-new",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.Modify:", resp)
	}

	{
		resp, err := devicetype.Delete("type-id")
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.Delete:", resp)
	}

	// device_type_point 示例
	{
		resp, err := devicetype.CreatePoint(devicetype.CreatePointRequest{
			DeviceTypeID: "type-id",
			PointID:      "point-1",
			Name:         "温度",
			Type:         "float",
			AccessMode:   "r",
			OrderNumber:  1,
			Description:  "温度点位",
			Group:        "环境",
			Unit:         "℃",
			Format:       "0.0",
			IsArray:      false,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.CreatePoint:", resp)
	}

	{
		resp, err := devicetype.UpdatePoint(devicetype.UpdatePointRequest{
			DeviceTypeID: "type-id",
			PointID:      "point-1",
			Name:         "温度(修改)",
			Type:         "float",
			AccessMode:   "rw",
			OrderNumber:  2,
			Description:  "温度点位-修改",
			Group:        "环境",
			Unit:         "℃",
			Format:       "0.0",
			IsArray:      false,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.UpdatePoint:", resp)
	}

	{
		resp, err := devicetype.DeletePoint("type-id", "point-1")
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.DeletePoint:", resp)
	}

	{
		resp, err := devicetype.DeleteAllPoints("type-id")
		if err != nil {
			panic(err)
		}
		fmt.Println("devicetype.DeleteAllPoints:", resp)
	}

	// device 示例
	{
		resp, err := device.Create(device.CreateRequest{
			Name:        "测试设备",
			TypeID:      "type-id",
			Description: "SDK创建测试设备",
			ConnectMode: "direct",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("device.Create:", resp)
	}

	{
		resp, err := device.Modify(device.ModifyRequest{
			ID:          "device-id",
			Name:        "新设备名",
			Description: "修改描述",
			ConnectMode: "gateway",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("device.Modify:", resp)
	}

	{
		resp, err := device.Delete("device-id")
		if err != nil {
			panic(err)
		}
		fmt.Println("device.Delete:", resp)
	}
}
