# TopStack Go SDK

TopStack Go SDK 是用于与 TopStack 平台 API 交互的 Go 语言客户端库，支持物联网、告警、能源、工单、全局变量等模块的接口调用。

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

### 物联网设备管理
```go
import "topstack-sdk-go/iot"

// 查询设备列表
devices, err := iot.QueryDeviceList(iot.DeviceListRequest{PageNum: 1, PageSize: 10})
if err != nil {
    // 错误处理
}
```

### 告警管理
```go
import "topstack-sdk-go/alert"

// 查询告警等级
levels, err := alert.QueryAlertLevel()
if err != nil {
    // 错误处理
}
```

### 能源管理
```go
import "topstack-sdk-go/ems"

// 查询仪表列表
meters, err := ems.QueryMeterList(ems.MeterListRequest{PageNum: 1, PageSize: 10})
if err != nil {
    // 错误处理
}
```

### 工单管理
```go
import "topstack-sdk-go/asset"

// 查询告警工单
orders, err := asset.QueryAlertWorkOrderList(asset.AlertWorkOrderListRequest{PageNum: 1, PageSize: 10})
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
- `iot/`    物联网设备、测点、网关等接口
- `alert/`  告警等级、类型、记录等接口
- `ems/`    能源管理相关接口
- `asset/`  工单相关接口
- `client/` HTTP 客户端封装、全局变量接口
- `sample/` 示例代码

## 错误处理
所有接口均返回标准 Go error：
```go
resp, err := iot.QueryDeviceList(req)
if err != nil {
    // 处理错误
}
```

## 配置说明
可通过代码或环境变量配置：
- `TOPSTACK_BASE_URL`
- `TOPSTACK_API_KEY`
- `TOPSTACK_SECRET`

## 示例项目
查看 `sample/` 目录下的完整示例：
```bash
cd sample
go run main.go
```

## 贡献
欢迎提交 Issue 和 Pull Request 改进 SDK。

## 许可证
[请补充您的许可证信息]

## 支持
- 邮箱: service@iotopo.com
- 文档: https://www.iotopo.com/docs/guide