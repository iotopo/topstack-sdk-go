package nats

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log/slog"
)

type Subscriber interface {
	Unsubscribe() error
}

type NatsConfig struct {
	Addr     string
	Token    string
	Username string
	Password string
}

type NatsBus struct {
	conn *nats.Conn
}

func (bus *NatsBus) Close() {
	if bus.conn != nil {
		bus.conn.Close()
	}
}

// SubscribePointData 订阅设备测点数据
func (bus *NatsBus) SubscribePointData(projectID, deviceID, pointID string, fn func(data *PointData)) (subscriber Subscriber, err error) {
	topic := realtimePointTopic(projectID, deviceID, pointID)
	return bus.conn.Subscribe(topic, func(msg *nats.Msg) {
		var data PointData
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal realtime point data error", "err", err)
			return
		}
		fn(&data)
	})
}

// SubscribeDeviceTypeData 订阅同设备模型下的测点数据
func (bus *NatsBus) SubscribeDeviceTypeData(projectID, deviceTypeID, pointID string, fn func(data *PointData)) (subscriber Subscriber, err error) {
	topic := realtimePointTopicV2(projectID, deviceTypeID, "*", pointID)
	return bus.conn.Subscribe(topic, func(msg *nats.Msg) {
		var data PointData
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal realtime point data error", "err", err)
			return
		}
		fn(&data)
	})
}

// SubscribeDeviceState 订阅设备状态数据
func (bus *NatsBus) SubscribeDeviceState(projectID, deviceID string, fn func(data *DeviceState)) (subscriber Subscriber, err error) {
	return bus.conn.Subscribe(deviceStateTopic(projectID, deviceID), func(msg *nats.Msg) {
		var data DeviceState
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal device state data error", "err", err)
			return
		}
		fn(&data)
	})
}

// SubscribeGatewayState 订阅网关状态数据
func (bus *NatsBus) SubscribeGatewayState(projectID string, fn func(topic string, data *GatewayState)) (subscriber Subscriber, err error) {
	return bus.conn.Subscribe(gatewayStateTopic(projectID, "*"), func(msg *nats.Msg) {
		var data GatewayState
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal gateway state data error", "err", err)
			return
		}
		fn(msg.Subject, &data)
	})
}

// SubscribeChannelState 订阅数据通道状态数据
func (bus *NatsBus) SubscribeChannelState(projectID string, fn func(topic string, data *ChannelState)) (subscriber Subscriber, err error) {
	return bus.conn.Subscribe(channelStateTopic(projectID, "*"), func(msg *nats.Msg) {
		var data ChannelState
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal channel state data error", "err", err)
			return
		}
		fn(msg.Subject, &data)
	})
}

// SubscribeAlertInfo 订阅全部告警消息
func (bus *NatsBus) SubscribeAlertInfo(projectID string, fn func(topic string, data *AlertInfo)) (subscriber Subscriber, err error) {
	return bus.conn.Subscribe(alertTopic(projectID), func(msg *nats.Msg) {
		var data AlertInfo
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal alert info data error", "err", err)
			return
		}
		fn(msg.Subject, &data)
	})
}

// SubscribeDeviceAlertInfo 订阅设备告警信息
func (bus *NatsBus) SubscribeDeviceAlertInfo(projectID, deviceID string, fn func(topic string, data *AlertInfo)) (subscriber Subscriber, err error) {
	return bus.conn.Subscribe(deviceAlertTopic(projectID, deviceID), func(msg *nats.Msg) {
		var data AlertInfo
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			slog.Error("unmarshal alert info data error", "err", err)
			return
		}
		fn(msg.Subject, &data)
	})
}

func NewNatsBus(conf NatsConfig, options ...nats.Option) (*NatsBus, error) {
	opts := []nats.Option{
		nats.MaxReconnects(-1),
		nats.RetryOnFailedConnect(true),
		//nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
		//	slog.Info("nats disconnected", "err", err)
		//}),
		//nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, err error) {
		//	slog.Error("nats error", "err", err)
		//}),
		//nats.ReconnectHandler(func(c *nats.Conn) {
		//	slog.Info("nats reconnected")
		//}),
		//nats.ConnectHandler(func(c *nats.Conn) {
		//	slog.Info("nats connected")
		//}),
		//nats.Name("myapp"),
	}

	token := conf.Token
	if token != "" {
		opts = append(opts, nats.Token(token))
	}

	username := conf.Username
	password := conf.Password
	if username != "" && password != "" {
		opts = append(opts, nats.UserInfo(username, password))
	}
	opts = append(opts, options...)

	nc, err := nats.Connect(conf.Addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("create nats connect error: %v", err)
	}
	return &NatsBus{conn: nc}, nil
}
