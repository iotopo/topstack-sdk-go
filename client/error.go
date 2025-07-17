package client

type HTTPError struct {
	StatusCode int
	Status     string
}

func (err *HTTPError) Error() string {
	return err.Status // fmt.Sprintf("%d %s", err.StatusCode, err.State)
}

func newHTTPError(statusCode int, status string) *HTTPError {
	return &HTTPError{StatusCode: statusCode, Status: status}
}

type ResponseError struct {
	Code string
	Msg  string
}
