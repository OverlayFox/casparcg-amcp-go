package types

import (
	"fmt"
	"strconv"
	"strings"
)

type BasicCommandInterface interface {
	String() string
}

type BasicCommand struct {
	VideoChannel int
	Layer        int
}

type CommandLoad struct {
	BasicCommand
	Clip string

	Parameters *map[string]string
}

func (c CommandLoad) String() string {
	cmd := fmt.Sprintf("LOAD %d-%d %s", c.VideoChannel, c.Layer, quote(c.Clip))
	if c.Parameters != nil {
		cmd += " " + buildParams(*c.Parameters)
	}
	return cmd
}

type CommandPlay struct {
	BasicCommand

	Clip       *string
	Parameters *map[string]string
}

func (c CommandPlay) String() string {
	cmd := fmt.Sprintf("PLAY %d-%d", c.VideoChannel, c.Layer)
	if c.Clip != nil {
		cmd += " " + quote(*c.Clip)
		if c.Parameters != nil {
			cmd += " " + buildParams(*c.Parameters)
		}
	}
	return cmd
}

type CommandPause struct {
	BasicCommand
}

func (c CommandPause) String() string {
	return fmt.Sprintf("PAUSE %d-%d", c.VideoChannel, c.Layer)
}

type CommandResume struct {
	BasicCommand
}

func (c CommandResume) String() string {
	return fmt.Sprintf("RESUME %d-%d", c.VideoChannel, c.Layer)
}

type CommandStop struct {
	BasicCommand
}

func (c CommandStop) String() string {
	return fmt.Sprintf("STOP %d-%d", c.VideoChannel, c.Layer)
}

type CommandClear struct {
	BasicCommand
}

func (c CommandClear) String() string {
	return fmt.Sprintf("CLEAR %d-%d", c.VideoChannel, c.Layer)
}

type CommandCall struct {
	BasicCommand
	Params map[string]string
}

func (c CommandCall) String() string {
	cmd := fmt.Sprintf("CALL %d-%d", c.VideoChannel, c.Layer)
	if len(c.Params) > 0 {
		cmd += " " + buildParams(c.Params)
	}
	return cmd
}

type CommandSwap struct {
	VideoChannel1 int
	Layer1        *int

	VideoChannel2 int
	Layer2        *int

	Transform bool // either nil or "TRANSFORMS"
}

func (c CommandSwap) String() string {
	cmd := "SWAP " + strconv.Itoa(c.VideoChannel1)
	if c.Layer1 != nil {
		cmd += "-" + strconv.Itoa(*c.Layer1)
	}
	cmd += " " + strconv.Itoa(c.VideoChannel2)
	if c.Layer2 != nil {
		cmd += "-" + strconv.Itoa(*c.Layer2)
	}
	if c.Transform {
		cmd += " TRANSFORMS"
	}
	return cmd
}

// CommandAdd adds a consumer to the specified video channel
// Different consumers require different parameters.
type CommandAdd struct {
	VideoChannel int
	ConsumerIdx  *int
	ConsumerName string
	Parameters   map[string]string
}

func (c CommandAdd) String() string {
	cmd := "ADD " + strconv.Itoa(c.VideoChannel)
	if c.ConsumerIdx != nil {
		cmd += "-" + strconv.Itoa(*c.ConsumerIdx)
	}
	cmd += " " + c.ConsumerName
	if len(c.Parameters) > 0 {
		cmd += " " + buildParams(c.Parameters)
	}
	return cmd
}

type CommandRemove struct {
	VideoChannel int

	ConsumerIdx *int
	Parameters  *map[string]string
}

func (c CommandRemove) String() string {
	cmd := "REMOVE " + strconv.Itoa(c.VideoChannel)
	if c.ConsumerIdx != nil {
		cmd += "-" + strconv.Itoa(*c.ConsumerIdx)
	}
	if c.Parameters != nil && len(*c.Parameters) > 0 {
		cmd += " " + buildParams(*c.Parameters)
	}
	return cmd
}

type CommandPrint struct {
	VideoChannel int
}

func (c CommandPrint) String() string {
	return "PRINT " + strconv.Itoa(c.VideoChannel)
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
	Level AMCPLogLevel
}

func (c CommandLogLevel) String() string {
	return "LOG LEVEL " + string(c.Level)
}

type SetVariable string

const (
	SetVariableMode          SetVariable = "MODE"
	SetVariableChannelLayout SetVariable = "CHANNEL_LAYOUT"
)

// CommandSet changes the value of a channel variable.
type CommandSet struct {
	VideoChannel int

	Variable SetVariable
	Value    string
}

func (c CommandSet) String() string {
	return fmt.Sprintf("SET %d %s %s", c.VideoChannel, c.Variable, c.Value)
}

type LockAction string

const (
	LockActionAcquire LockAction = "ACQUIRE"
	LockActionRelease LockAction = "RELEASE"
	LockActionClear   LockAction = "CLEAR"
)

type CommandLock struct {
	VideoChannel int
	Action       LockAction
	Secret       *string
}

func (c CommandLock) String() string {
	cmd := fmt.Sprintf("LOCK %d %s", c.VideoChannel, c.Action)
	if c.Secret != nil {
		cmd += " " + *c.Secret
	}
	return cmd
}

type CommandPing struct {
	Token string
}

func (c CommandPing) String() string {
	if c.Token != "" {
		return "PING " + c.Token
	}
	return "PING"
}

type CommandBye struct {
}

func (c CommandBye) String() string {
	return "BYE"
}

type CommandKill struct {
}

func (c CommandKill) String() string {
	return "KILL"
}

type CommandRestart struct {
}

func (c CommandRestart) String() string {
	return "RESTART"
}

// Helper functions for command serialization

func quote(s string) string {
	return "\"" + s + "\""
}

func buildParams(params map[string]string) string {
	var parts []string
	for k, v := range params {
		parts = append(parts, k+" "+v)
	}
	return strings.Join(parts, " ")
}
