package common

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/iotopo/topstack-sdk-go/topstack/common/errors"
	tchttp "github.com/iotopo/topstack-sdk-go/topstack/common/http"
)

const (
	codeLimitExceeded = "RequestLimitExceeded"
	tplRateLimitRetry = "[WARN] rate limit exceeded, retrying (%d/%d) in %f seconds: %s"
)

func (c *Client) sendWithRateLimitRetry(req *http.Request, retryable bool) (resp *http.Response, err error) {
	// make sure maxRetries is more than 0
	maxRetries := maxInt(c.profile.RateLimitExceededMaxRetries, 0)
	durationFunc := safeDurationFunc(c.profile.RateLimitExceededRetryDuration)

	var shadow []byte
	for idx := 0; idx <= maxRetries; idx++ {
		resp, err = c.sendWithNetworkFailureRetry(req, retryable)
		if err != nil {
			return
		}

		resp.Body, shadow = shadowRead(resp.Body)

		err = tchttp.ParseErrorFromHTTPResponse(shadow)
		// should not sleep on last request
		if err, ok := err.(*errors.TopStackSDKError); ok && err.Code == codeLimitExceeded && idx < maxRetries {
			duration := durationFunc(idx)
			if c.debug {
				log.Printf(tplRateLimitRetry, idx, maxRetries, duration.Seconds(), err.Error())
			}

			time.Sleep(duration)
			continue
		}

		return resp, err
	}

	return resp, err
}

func shadowRead(reader io.ReadCloser) (io.ReadCloser, []byte) {
	val, err := ioutil.ReadAll(reader)
	if err != nil {
		return reader, nil
	}
	return ioutil.NopCloser(bytes.NewBuffer(val)), val
}
