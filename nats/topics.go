package nats

import "fmt"

func realtimePointTopicV2(projectID, deviceTypeID, deviceID, pointID string) string {
	return fmt.Sprintf("iot.platform.device.datas.%s.%s.%s.%s", projectID, deviceTypeID, deviceID, pointID)
}

func realtimePointTopic(projectID, deviceID, pointID string) string {
	return realtimePointTopicV2(projectID, "*", deviceID, pointID)
}

func channelStateTopic(projectID, channelID string) string {
	return fmt.Sprintf("iot.platform.channel.state.%s.%s", projectID, channelID)
}
func gatewayStateTopic(projectID, gatewayID string) string {
	return fmt.Sprintf("iot.platform.gateway.state.%s.%s", projectID, gatewayID)
}

func deviceStateTopic(projectID, deviceID string) string {
	return fmt.Sprintf("iot.platform.device.state.%s.%s", projectID, deviceID)
}

func alertTopic(projectID string) string {
	return fmt.Sprintf("iot.platform.alert.%s.>", projectID)
}

func deviceAlertTopic(projectID, deviceID string) string {
	if deviceID == "" {
		deviceID = "not_device"
	}
	//return fmt.Sprintf("iot.platform.alert.%s.%s.%s", projectID, deviceTypeID, deviceID)
	return fmt.Sprintf("iot.platform.alert.%s.%s", projectID, deviceID)
}
