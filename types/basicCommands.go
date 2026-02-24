package types

type BasicCommandInterface interface {
}

type BasicCommand struct {
	videoChannel int
	layer        int
}

type CommandLoad struct {
	BasicCommand
	clip string

	parameters *map[string]string
}

type CommandPlay struct {
	BasicCommand

	clip       *string
	parameters *map[string]string
}

type CommandPause struct {
	BasicCommand
}

type CommandResume struct {
	BasicCommand
}

type CommandStop struct {
	BasicCommand
}

type CommandClear struct {
	BasicCommand
}

type CommandCall struct {
	BasicCommand
	params map[string]string
}

type CommandSwap struct {
	videoChannel1 int
	layer1        *int

	videoChannel2 int
	layer2        *int

	transform bool // either nil or "TRANSFORMS"
}

// CommandAdd adds a consumer to the specified video channel
// Different consumers require different parameters.
type CommandAdd struct {
	videoChannel int
	consumerIdx  *int
	consumerName string
	parameters   map[string]string
}

type CommandRemove struct {
	videoChannel int

	consumerIdx *int
	parameters  *map[string]string
}

type CommandPrint struct {
	videoChannel int
}

type AMCPLogLevel string

const (
	AMCPLogLevelTrace AMCPLogLevel = "trace"
	AMCPLogLevelDebug AMCPLogLevel = "debug"
	AMCPLogLevelInfo  AMCPLogLevel = "info"
	AMCPLogLevelWarn  AMCPLogLevel = "warn"
	AMCPLogLevelError AMCPLogLevel = "error"
	AMCPLogLevelFatal AMCPLogLevel = "fatal"
)

type CommandLogLevel struct {
	level AMCPLogLevel
}

type SetVariable string

const (
	SetVariableMode          SetVariable = "MODE"
	SetVariableChannelLayout SetVariable = "CHANNEL_LAYOUT"
)

// CommandSet changes the value of a channel variable.
type CommandSet struct {
	videoChannel int

	variable SetVariable
	value    string
}

type LockAction string

const (
	LockActionAcquire LockAction = "ACQUIRE"
	LockActionRelease LockAction = "RELEASE"
	LockActionClear   LockAction = "CLEAR"
)

type CommandLock struct {
	videoChannel int
	action       LockAction
	secret       *string
}

type CommandPing struct {
	token string
}

type CommandBye struct {
}

type CommandKill struct {
}

type CommandRestart struct {
}
