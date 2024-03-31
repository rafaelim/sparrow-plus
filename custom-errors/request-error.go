package customerrors

import "fmt"

type RequestError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Reason     string `json:"reason"`
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("%v; Reason: %v", r.Message, r.Reason)
}

func NewRequestError(statusCode int, message, reason string) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Message:    message,
		Reason:     reason,
	}
}

func (e *RequestError) ResponseHeaders() (int, map[string]string) {
	return e.StatusCode, map[string]string{
		"Content-Type": "application/json",
	}
}
