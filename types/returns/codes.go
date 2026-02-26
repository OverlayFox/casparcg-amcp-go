package returns

// ReturnCode is a type that represents the return code of an operation or function. It is typically used to indicate the success or failure of an operation, as well as to provide additional information about the outcome. The specific values and meanings of return codes can vary depending on the context in which they are used.
type ReturnCode int

const (
	InformationEvent         ReturnCode = 100 // Information about an event.
	InformationEventWithData ReturnCode = 101 // Information about an event. A line of data is being returned.
)

const (
	SuccessWithMultilineData  ReturnCode = 200 // Success. Multiple lines of data are being returned.
	SuccessWithSinglelineData ReturnCode = 201 // Success. A single line of data is being returned.
	Success                   ReturnCode = 202 // Success. No data is being returned.
)

const (
	FailureNotUnderstood       ReturnCode = 400 // Failure. The command was not understood.
	FailureIllegalVideoChannel ReturnCode = 401 // Failure. The specified video channel is illegal.
	FailureParameterMissing    ReturnCode = 402 // Failure. A required parameter is missing.
	FailureIllegalParameter    ReturnCode = 403 // Failure. An illegal parameter value was provided.
	FailureMediaNotFound       ReturnCode = 404 // Failure. The specified media was not found.
)

const (
	ServerInternalError         ReturnCode = 500 // Server internal error.
	ServerInternalErrorWithData ReturnCode = 501 // Server internal error. A line of data is being returned.
	ServerMediaFileUnreadable   ReturnCode = 502 // Server internal error. The media file is unreadable.
	ServerAccessError           ReturnCode = 503 // Server internal error. Access to the media file is denied.
)
