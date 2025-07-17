package alertrecord

import "topstack-sdk-go/client"

// Activity 活动告警查询
// GET /alert/open_api/v1/alert_record/activity
func Activity(req ActivityRequest) (resp client.Response[ActivityResponse], err error) {
	err = client.GET("/alert/open_api/v1/alert_record/activity", req, &resp)
	return
}

// IgnoreBatch 告警记录批量忽略
// PUT /alert/open_api/v1/alert_record/ignoredBatch
func IgnoreBatch(req IgnoredBatchRequest) (resp client.Response[IgnoredBatchResponse], err error) {
	err = client.PUT("/alert/open_api/v1/alert_record/ignoredBatch", req, &resp)
	return
}
