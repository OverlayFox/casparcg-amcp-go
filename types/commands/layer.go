package commands

import (
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type LayerCommand struct {
	VideoChannel int
	Layer        *int // Layer is optional for channel-level commands
}

type LayerLoad struct {
	LayerCommand

	Clip string

	Parameters *[]string
}

func (c LayerLoad) String() string {
	cmd := baseCommand("LOAD", c.VideoChannel, c.Layer)
	cmd = appendQuotedString(cmd, &c.Clip)
	cmd = appendParams(cmd, c.Parameters)
	return cmd
}

type LayerPlay struct {
	LayerCommand

	Clip       *string
	Parameters *[]string
}

func (c LayerPlay) String() string {
	cmd := baseCommand("PLAY", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, c.Clip)
	cmd = appendParams(cmd, c.Parameters)
	return cmd
}

type LayerPause struct {
	LayerCommand
}

func (c LayerPause) String() string {
	return baseCommand("PAUSE", c.VideoChannel, c.Layer)
}

type LayerResume struct {
	LayerCommand
}

func (c LayerResume) String() string {
	return baseCommand("RESUME", c.VideoChannel, c.Layer)
}

type LayerStop struct {
	LayerCommand
}

func (c LayerStop) String() string {
	return baseCommand("STOP", c.VideoChannel, c.Layer)
}

type LayerClear struct {
	LayerCommand
}

func (c LayerClear) String() string {
	return baseCommand("CLEAR", c.VideoChannel, c.Layer)
}

type LayerCall struct {
	LayerCommand

	Params []string
}

func (c LayerCall) String() string {
	cmd := baseCommand("CALL", c.VideoChannel, c.Layer)
	cmd = appendParams(cmd, &c.Params)
	return cmd
}

type LayerSwap struct {
	LayerCommand

	VideoChannel2 int
	Layer2        *int

	Transform bool // either nil or "TRANSFORMS"
}

func (c LayerSwap) String() string {
	cmd := baseCommand("SWAP", c.VideoChannel, c.Layer)
	cmd += " " + strconv.Itoa(c.VideoChannel2)
	if c.Layer2 != nil {
		cmd += "-" + strconv.Itoa(*c.Layer2)
	}
	if c.Transform {
		cmd += " TRANSFORMS"
	}
	return cmd
}

type LayerAdd struct {
	LayerCommand

	ConsumerIdx  *int
	ConsumerName string // TODO: Make this an enum of possible consumer types
	Params       *[]string
}

func (c LayerAdd) String() string {
	cmd := baseCommand("ADD", c.VideoChannel, nil)
	if c.ConsumerIdx != nil {
		cmd += "-" + strconv.Itoa(*c.ConsumerIdx)
	}
	cmd = appendString(cmd, &c.ConsumerName)
	cmd = appendParams(cmd, c.Params)
	return cmd
}

type LayerRemove struct {
	LayerCommand

	ConsumerIdx *int
	Parameters  *[]string
}

func (c LayerRemove) String() string {
	cmd := baseCommand("REMOVE", c.VideoChannel, nil)
	if c.ConsumerIdx != nil {
		cmd += "-" + strconv.Itoa(*c.ConsumerIdx)
	} else {
		cmd = appendParams(cmd, c.Parameters)
	}
	return cmd
}

type LayerPrint struct {
	LayerCommand
}

func (c LayerPrint) String() string {
	return baseCommand("PRINT", c.VideoChannel, nil)
}

type LayerSet struct {
	LayerCommand

	VariableName types.SetVariable
	Value        string
}

func (c LayerSet) String() string {
	cmd := baseCommand("SET", c.VideoChannel, nil)
	cmd = appendString(cmd, ptr(c.VariableName.String()))
	cmd = appendString(cmd, &c.Value)
	return cmd
}

type LayerLock struct {
	LayerCommand

	Action     types.LockAction
	Passphrase *string
}

func (c LayerLock) String() string {
	cmd := baseCommand("LOCK", c.VideoChannel, nil)
	cmd = appendString(cmd, ptr(c.Action.String()))
	cmd = appendString(cmd, c.Passphrase)
	return cmd
}

type LayerInfo struct {
	LayerCommand
}

func (c LayerInfo) String() string {
	return baseCommand("INFO", c.VideoChannel, c.Layer)
}

type LayerInfoDelay struct {
	LayerCommand
}

func (c LayerInfoDelay) String() string {
	cmd := baseCommand("INFO", c.VideoChannel, c.Layer)
	cmd += " DELAY"
	return cmd
}
