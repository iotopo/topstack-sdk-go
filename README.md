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
- **实时数据订阅 (nats/)**：设备测点数据、设备状态、网关状态、数据通道状态、告警信息等
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
	//cli := client.Init("https://your-topstack-instance.com", apiKey, projectID)
    cli := client.InitClient("https://your-topstack-instance.com", appKey, appSecret)
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

### NATS 实时数据订阅

TopStack SDK 支持通过 NATS 消息总线订阅实时数据，包括设备测点数据、设备状态、网关状态、数据通道状态和告警信息。

#### 初始化 NATS 连接
```go
import "topstack-sdk-go/nats"

// 配置 NATS 连接参数
config := nats.NatsConfig{
    Addr:     "nats://your-nats-server:4222",
    Token:    "your-nats-token",        // 可选
    Username: "your-username",          // 可选
    Password: "your-password",          // 可选
}

// 创建 NATS 总线
bus, err := nats.NewNatsBus(config)
if err != nil {
    // 错误处理
}
defer bus.Close()
```

#### 订阅设备测点数据
```go
// 订阅单个设备的测点数据
subscriber, err := bus.SubscribePointData("project1", "device1", "point1", func(data *nats.PointData) {
    fmt.Printf("收到测点数据: 设备=%s, 测点=%s, 值=%v, 质量=%d, 时间=%v\n",
        data.DeviceID, data.PointID, data.Value, data.Quality, data.Timestamp)
})
if err != nil {
    // 错误处理
}
defer subscriber.Unsubscribe()

// 订阅同设备模型下的所有设备测点数据
subscriber, err := bus.SubscribeDeviceTypeData("project1", "deviceType1", "point1", func(data *nats.PointData) {
    fmt.Printf("收到设备模型测点数据: 设备=%s, 测点=%s, 值=%v\n",
        data.DeviceID, data.PointID, data.Value)
})
if err != nil {
    // 错误处理
}
defer subscriber.Unsubscribe()
```

#### 订阅设备状态
```go
// 订阅设备在线/离线状态
subscriber, err := bus.SubscribeDeviceState("project1", "device1", func(data *nats.DeviceState) {
    status := "离线"
    if data.State == 1 {
        status = "在线"
    }
    fmt.Printf("设备状态变化: 设备=%s, 状态=%s, 时间=%v\n",
        data.DeviceID, status, data.Timestamp)
})
if err != nil {
    // 错误处理
}
defer subscriber.Unsubscribe()
```

#### 订阅网关状态
```go
// 订阅所有网关状态（使用通配符）
subscriber, err := bus.SubscribeGatewayState("project1", func(topic string, data *nats.GatewayState) {
    status := "离线"
    if data.State == 1 {
        status = "在线"
    }
    fmt.Printf("网关状态变化: 网关=%s, 状态=%s, 时间=%v\n",
        data.Name, status, data.Timestamp)
})
if err != nil {
    // 错误处理
}
defer subscriber.Unsubscribe()
```

#### 订阅数据通道状态
```go
// 订阅所有数据通道状态
subscriber, err := bus.SubscribeChannelState("project1", func(topic string, data *nats.ChannelState) {
    fmt.Printf("数据通道状态: 通道=%s, 运行=%v, 连接=%v, 时间=%v\n",
        data.ChannelName, data.Running, data.Connected, data.Timestamp)
})
if err != nil {
    // 错误处理
}
defer subscriber.Unsubscribe()
```

#### 订阅告警信息
```go
// 订阅项目下的所有告警信息
subscriber, err := bus.SubscribeAlertInfo("project1", func(topic string, data *nats.AlertInfo) {
    fmt.Printf("收到告警: 标题=%s, 等级=%s, 设备=%s, 时间=%v\n",
        data.Title, data.AlertLevelName, data.DeviceName, data.CreatedAt)
})
if err != nil {
    // 错误处理
}
defer subscriber.Unsubscribe()
```

#### 数据模型说明

**PointData（测点数据）**
- `DeviceID`: 设备ID
- `PointID`: 测点ID
- `Value`: 测点值
- `Quality`: 数据质量（1=离线，2=无效）
- `Timestamp`: 时间戳
- `Status`: 状态（0=正常，>0=越上限，<0=越下限）

**DeviceState（设备状态）**
- `DeviceID`: 设备ID
- `State`: 状态（0=离线，1=在线）
- `Timestamp`: 时间戳

**GatewayState（网关状态）**
- `Name`: 网关名称
- `GatewayID`: 网关ID
- `State`: 状态（0=离线，1=在线）
- `Timestamp`: 时间戳

**ChannelState（数据通道状态）**
- `ChannelID`: 通道ID
- `ChannelName`: 通道名称
- `Running`: 是否运行中
- `Connected`: 是否已连接
- `Timestamp`: 时间戳

**AlertInfo（告警信息）**
- `Title`: 告警标题
- `Content`: 告警内容
- `AlertLevelName`: 告警等级名称
- `DeviceName`: 设备名称
- `CreatedAt`: 创建时间
- `Status`: 告警状态（unhandled/handled/ignored/auto）

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
- `nats/`   NATS 实时数据订阅接口
  - `nats/nats.go`    NATS 总线实现
  - `nats/model.go`   数据模型定义
  - `nats/topics.go`  主题格式定义
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