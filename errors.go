package casparcg

import "errors"

var (
	ErrValueOutOfRange = errors.New("value out of range")
)

type CasparCGError struct {
	Code    int
	Message string
}

func (e CasparCGError) Error() string {
	return e.Message
}
