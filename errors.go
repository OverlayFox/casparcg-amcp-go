package casparcg

import "errors"

var (
	ErrValueOutOfRange = errors.New("value out of range")
	ErrLayerNotSet     = errors.New("layer is not set")
	ErrLayerSet        = errors.New("layer is set but expected to be unset")
)

type CasparCGError struct {
	Code    int
	Message string
}

func (e CasparCGError) Error() string {
	return e.Message
}
