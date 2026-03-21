package types

type CasparCGError struct {
	Code    int
	Message string
}

func (e CasparCGError) Error() string {
	return e.Message
}
