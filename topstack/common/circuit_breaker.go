package common

import (
	"github.com/iotopo/topstack-sdk-go/topstack/common/gobreaker"
	"time"
)

func defaultBreaker() *gobreaker.CircuitBreaker {
	defaultSet := gobreaker.Settings {
		MaxRequests: 5,
		Interval: 1 * 60 * time.Second,
		Timeout: 60 * time.Second,
	}
	return gobreaker.NewCircuitBreaker(defaultSet)
}
