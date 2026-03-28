package commands

import (
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type LayerCommand struct {
	VideoChannel int
	Layer        *int // Layer is optional for channel-level commands
}

type LayerCommandLoad struct {
	LayerCommand

	Clip string

	Parameters *[]string
}

func (c LayerCommandLoad) String() string {
	cmd := baseLayerCommand("LOAD", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, &c.Clip)
	cmd = appendParams(cmd, c.Parameters)
	return cmd
}

type LayerCommandPlay struct {
	LayerCommand

	Clip       *string
	Parameters *[]string
}

func (c LayerCommandPlay) String() string {
	cmd := baseLayerCommand("PLAY", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, c.Clip)
	cmd = appendParams(cmd, c.Parameters)
	return cmd
}

type LayerCommandPause struct {
	LayerCommand
}

func (c LayerCommandPause) String() string {
	return baseLayerCommand("PAUSE", c.VideoChannel, c.Layer)
}

type LayerCommandResume struct {
	LayerCommand
}

func (c LayerCommandResume) String() string {
	return baseLayerCommand("RESUME", c.VideoChannel, c.Layer)
}

type LayerCommandStop struct {
	LayerCommand
}

func (c LayerCommandStop) String() string {
	return baseLayerCommand("STOP", c.VideoChannel, c.Layer)
}

type LayerCommandClear struct {
	LayerCommand
}

func (c LayerCommandClear) String() string {
	return baseLayerCommand("CLEAR", c.VideoChannel, c.Layer)
}

type LayerCommandCall struct {
	LayerCommand

	Params []string
}

func (c LayerCommandCall) String() string {
	cmd := baseLayerCommand("CALL", c.VideoChannel, c.Layer)
	cmd = appendParams(cmd, &c.Params)
	return cmd
}

type LayerCommandSwap struct {
	LayerCommand

	VideoChannel2 int
	Layer2        *int

	Transform bool // either nil or "TRANSFORMS"
}

func (c LayerCommandSwap) String() string {
	cmd := baseLayerCommand("SWAP", c.VideoChannel, c.Layer)
	cmd += " " + strconv.Itoa(c.VideoChannel2)
	if c.Layer2 != nil {
		cmd += "-" + strconv.Itoa(*c.Layer2)
	}
	if c.Transform {
		cmd += " TRANSFORMS"
	}
	return cmd
}

type LayerCommandAdd struct {
	LayerCommand

	ConsumerIdx  *int      // ConsumerIdx overrides the index that the consumer itself decides and can later be used with the REMOVE command to remove the consumer.
	ConsumerName string    // TODO: Make this an enum of possible consumer types
	Parameters   *[]string // Parameters are specific to the consumer being added. For example, for a STREAM consumer you can add []string{"udp://localhost:5004", "-vcodec", "libx264", "-tune", "zerolatency"}
}

func (c LayerCommandAdd) String() string {
	cmd := baseLayerCommand("ADD", c.VideoChannel, nil)
	if c.ConsumerIdx != nil {
		cmd += "-" + strconv.Itoa(*c.ConsumerIdx)
	}
	appendString(cmd, &c.ConsumerName)
	appendParams(cmd, c.Parameters)
	return cmd
}

type LayerCommandRemove struct {
	LayerCommand

	ConsumerIdx *int      // ConsumerIdx overrides the index that the consumer itself decides and can later be used with the REMOVE command to remove the consumer.
	Parameters  *[]string // Parameters are specific to the consumer being added. For example, for a STREAM consumer you can add []string{"udp://localhost:5004", "-vcodec", "libx264", "-tune", "zerolatency"}
}

func (c LayerCommandRemove) String() string {
	cmd := baseLayerCommand("REMOVE", c.VideoChannel, nil)
	if c.ConsumerIdx != nil {
		cmd += "-" + strconv.Itoa(*c.ConsumerIdx)
	} else {
		appendParams(cmd, c.Parameters)
	}
	return cmd
}

type LayerCommandPrint struct {
	LayerCommand
}

func (c LayerCommandPrint) String() string {
	return baseLayerCommand("PRINT", c.VideoChannel, nil)
}

type LayerCommandSet struct {
	LayerCommand

	VariableName string
	Value        string
}

func (c LayerCommandSet) String() string {
	cmd := baseLayerCommand("SET", c.VideoChannel, nil)
	cmd = appendString(cmd, &c.VariableName)
	cmd = appendString(cmd, &c.Value)
	return cmd
}

type LayerCommandLock struct {
	LayerCommand

	Action     types.LockAction
	Passphrase *string
}

func (c LayerCommandLock) String() string {
	cmd := baseLayerCommand("LOCK", c.VideoChannel, nil)
	action := c.Action.String()
	cmd = appendString(cmd, &action)
	cmd = appendString(cmd, c.Passphrase)
	return cmd
}

type LayerCommandInfo struct {
	LayerCommand
}

func (c LayerCommandInfo) String() string {
	return baseLayerCommand("INFO", c.VideoChannel, c.Layer)
}
