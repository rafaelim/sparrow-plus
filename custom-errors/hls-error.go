package customerrors

import "fmt"

type HLSError struct {
	Message string
	Reason  string
}

func (r *HLSError) Error() string {
	return fmt.Sprintf("message: %v reason: %v", r.Message, r.Reason)
}

func NewHLSError(message, reason string) *HLSError {
	return &HLSError{
		Message: message,
		Reason:  reason,
	}
}
