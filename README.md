# TopStack Go SDK

官方网址：https://www.iotopo.com

TopStack 是一款轻量型 Web 组态软件，提供设备数据采集、定时任务、控制策略、联动控制、设备告警、设备维护管理、设备绩效管理、能源管理、组态开发、报表开发等功能。支持移动端访问，支持本地部署，可帮助企业从无到有快速搭建工业物联网平台。

TopStack 目前已完成了信创生态的全面适配：

* 适配国产服务器：龙芯、飞腾、鲲鹏、海光、兆芯等。
* 适配国产操作系统：麒麟、统信等操作系统。
* 适配国产数据库：达梦、人大金仓等国产数据库。

本项目为 TopStack SDK Go 版本，是用于与 TopStack 平台交互的 Go 语言客户端库，支持物联网、告警、能源、工单、全局变量等模块的接口调用。


## 功能特性
- **物联网管理 (iot/)**：设备、测点、网关、数据等
- **告警管理 (alert/)**：告警等级、类型、记录等
- **能源管理 (ems/)**：仪表、能源类型、用能单元、分项、能耗报表等
- **工单管理 (asset/)**：告警工单、现场工单、维护工单、计划工单等
- **全局变量 (client/)**：全局变量的读取与更新

## 安装
```bash
go get github.com/iotopo/topstack-sdk-go
```

## 快速开始

### 初始化客户端
```go
import (
    "topstack-sdk-go/client"
)

func main() {
	apiKey := "your-api-key"
	projectID := "iotopo"
	cli := client.Init("https://your-topstack-instance.com", apiKey, projectID)
	cli.SetDebug(true)
    // ...
}
```

### 物联网数据操作
```go
import "topstack-sdk-go/iot"

// 查询单测点实时值
resp, err := iot.FindLast(iot.FindLastRequest{DeviceID: "dev1", PointID: "v1"})
if err != nil {
    // 错误处理
}

// 批量查询多测点实时值
resp, err := iot.FindLastBatch(iot.FindLastBatchRequest{{DeviceID: "dev1", PointID: "v1"}})
if err != nil {
    // 错误处理
}

// 控制指令下发
resp, err := iot.SetValue(iot.SetValueRequest{DeviceID: "dev1", PointID: "V4", Value: "2"})
if err != nil {
    // 错误处理
}

// 查询历史数据
resp, err := iot.QueryHistory(iot.HistoryRequest{
    Points:      []iot.FindLastRequest{{DeviceID: "dev1", PointID: "v1"}},
    Start:       time.Now().Add(-10 * time.Minute),
    End:         time.Now(),
    Interval:    "10s",
    Aggregation: "last",
})
if err != nil {
    // 错误处理
}
```

### 物联网设备管理
```go
import "topstack-sdk-go/iot/device"

// 查询设备列表
resp, err := device.Query(device.QueryRequest{Search: "dev1", PageNum: 1, PageSize: 10})
if err != nil {
    // 错误处理
}

// 查询设备属性
resp, err := device.QueryProps("dev1")
if err != nil {
    // 错误处理
}

// 查询设备测点
resp, err := device.QueryPoint(device.PointQueryRequest{DeviceID: "dev1"})
if err != nil {
    // 错误处理
}
```

### 物联网设备类型管理
```go
import "topstack-sdk-go/iot/devicetype"

// 查询设备类型列表
resp, err := devicetype.Query(devicetype.QueryRequest{Search: "type1"})
if err != nil {
    // 错误处理
}

// 查询设备类型测点
resp, err := devicetype.QueryPoint(devicetype.PointQueryRequest{DeviceTypeID: "type1"})
if err != nil {
    // 错误处理
}
```

### 物联网网关管理
```go
import "topstack-sdk-go/iot/gateway"

// 查询网关列表
resp, err := gateway.Query(gateway.QueryRequest{Search: "gateway1", PageNum: 1, PageSize: 10})
if err != nil {
    // 错误处理
}
```

### 告警管理
```go
import "topstack-sdk-go/alert/alertlevel"

// 查询告警等级
resp, err := alertlevel.Query()
if err != nil {
    // 错误处理
}
```

### 能源管理
```go
import "topstack-sdk-go/ems/meter"

// 查询仪表列表
resp, err := meter.Query(meter.QueryRequest{PageNum: 1, PageSize: 10})
if err != nil {
    // 错误处理
}
```

### 工单管理
```go
import "topstack-sdk-go/asset/workorder/alert"

// 查询告警工单
resp, err := alert.Query(alert.QueryRequest{PageNum: 1, PageSize: 10})
if err != nil {
    // 错误处理
}
```

### 全局变量
```go
import "topstack-sdk-go/client"

// 读取全局变量
val, err := client.GetGlobalVarValue(client.GlobalVarGetValueRequest{Namespace: "system", Name: "config"})
if err != nil {
    // 错误处理
}
```

## 目录结构与模块说明
- `iot/`    物联网数据操作接口
  - `iot/device/`    设备管理接口
  - `iot/devicetype/` 设备类型管理接口
  - `iot/gateway/`   网关管理接口
- `alert/`  告警管理接口
  - `alert/alertlevel/` 告警等级管理
  - `alert/alerttype/`  告警类型管理
  - `alert/alertrecord/` 告警记录管理
- `ems/`    能源管理相关接口
  - `ems/meter/`     仪表管理
  - `ems/sector/`    用能单元管理
  - `ems/subentry/`  分项管理
- `asset/`  工单相关接口
  - `asset/workorder/alert/`      告警工单
  - `asset/workorder/locale/`     现场工单
  - `asset/workorder/maintenance/` 维护工单
  - `asset/workorder/schedule/`   计划工单
- `client/` HTTP 客户端封装、全局变量接口

## 错误处理
所有接口均返回标准 Go error：
```go
resp, err := iot.FindLast(req)
if err != nil {
    // 处理错误
}
```

## 示例项目
查看 `sample/` 目录下的完整示例：
```bash
cd sample
go run main.go
```

## 贡献
欢迎提交 Issue 和 Pull Request 改进 SDK。

## 许可证
MIT License 

## 支持
- 邮箱: service@iotopo.com
- 文档: https://www.iotopo.com/docs/guide