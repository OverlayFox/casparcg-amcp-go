package casparcg

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
	"github.com/overlayfox/casparcg-amcp-go/types/responses"
)

// LayerBuilder provides a fluent interface for building layer-based commands.
type LayerBuilder struct {
	client *Client
}

// sendCommand abstracts sending a command that does not expect a response value.
func (b *LayerBuilder) sendCommand(cmd interface{ String() string }) error {
	_, err := b.client.Send(cmd)
	return err
}

// Layer creates a new layer command builder for the specified channel and layer.
func (c *Client) Layer() *LayerBuilder {
	return &LayerBuilder{
		client: c,
	}
}

//
// Channel Commands
//

type LayerChannelBuilder struct {
	LayerBuilder

	videoChannel int
}

func (c *LayerBuilder) Channel(videoChannel int) *LayerChannelBuilder {
	return &LayerChannelBuilder{
		LayerBuilder: *c,
		videoChannel: videoChannel,
	}
}

func (b *LayerChannelBuilder) baseLayerChannelCommand() commands.LayerCommand {
	return commands.LayerCommand{
		VideoChannel: b.videoChannel,
		Layer:        nil, // Layer is nil for channel-level commands
	}
}

// Clear clears all layers in the channel.
func (b *LayerChannelBuilder) Clear() error {
	cmd := commands.LayerCommandClear{
		LayerCommand: b.baseLayerChannelCommand(),
	}
	return b.sendCommand(cmd)
}

// Swap swaps channels.
func (b *LayerChannelBuilder) Swap(channel2 int, transforms bool) error {
	cmd := commands.LayerCommandSwap{
		LayerCommand:  b.baseLayerChannelCommand(),
		VideoChannel2: channel2,
		Layer2:        nil,
		Transform:     transforms,
	}
	return b.sendCommand(cmd)
}

func (b *LayerChannelBuilder) Add(params types.LayerAdd) error {
	cmd := commands.LayerCommandAdd{
		LayerCommand: b.baseLayerChannelCommand(),
		ConsumerName: params.ConsumerName,
		ConsumerIdx:  params.ConsumerIdx,
		Params:       params.Parameters,
	}
	return b.sendCommand(cmd)
}

type LayerCommandRemove struct {
	LayerChannelBuilder
}

// Remove removes an existing consumer from video_channel.
func (b *LayerChannelBuilder) Remove() *LayerCommandRemove {
	return &LayerCommandRemove{
		LayerChannelBuilder: *b,
	}
}

// ConsumerIDX removes the consumer via its id.
func (b *LayerCommandRemove) ConsumerIDX(idx int) error {
	cmd := commands.LayerCommandRemove{
		LayerCommand: b.baseLayerChannelCommand(),
		ConsumerIdx:  &idx,
		Parameters:   nil,
	}
	return b.sendCommand(cmd)
}

// Params removes the consumer that matches the given parameters.
func (b *LayerCommandRemove) Params(params []string) error {
	cmd := commands.LayerCommandRemove{
		LayerCommand: b.baseLayerChannelCommand(),
		ConsumerIdx:  nil,
		Parameters:   &params,
	}
	return b.sendCommand(cmd)
}

func (b *LayerChannelBuilder) Print() error {
	cmd := commands.LayerCommandPrint{
		LayerCommand: b.baseLayerChannelCommand(),
	}
	return b.sendCommand(cmd)
}

type LayerCommandSet struct {
	LayerChannelBuilder
}

func (b *LayerChannelBuilder) Set() *LayerCommandSet {
	return &LayerCommandSet{
		LayerChannelBuilder: *b,
	}
}

func (b *LayerCommandSet) Mode(value types.VideoMode) error {
	cmd := commands.LayerCommandSet{
		LayerCommand: b.baseLayerChannelCommand(),
		VariableName: types.SetVariableMode,
		Value:        string(value),
	}
	return b.sendCommand(cmd)
}

func (b *LayerCommandSet) ChannelLayout(value types.AudioChannelLayout) error {
	cmd := commands.LayerCommandSet{
		LayerCommand: b.baseLayerChannelCommand(),
		VariableName: types.SetVariableChannelLayout,
		Value:        string(value),
	}
	return b.sendCommand(cmd)
}

type LayerCommandLock struct {
	LayerChannelBuilder
}

// Lock locks a channel for exclusive access.
func (b *LayerChannelBuilder) Lock() *LayerCommandLock {
	return &LayerCommandLock{
		LayerChannelBuilder: *b,
	}
}

// Acquire acquires the lock with the given passphrase.
func (b *LayerCommandLock) Acquire(passphrase string) error {
	cmd := commands.LayerCommandLock{
		LayerCommand: b.baseLayerChannelCommand(),
		Action:       types.LockActionAcquire,
		Passphrase:   &passphrase,
	}
	return b.sendCommand(cmd)
}

// Release releases the lock with the given passphrase.
func (b *LayerCommandLock) Release(passphrase string) error {
	cmd := commands.LayerCommandLock{
		LayerCommand: b.baseLayerChannelCommand(),
		Action:       types.LockActionRelease,
		Passphrase:   &passphrase,
	}
	return b.sendCommand(cmd)
}

// Clear clears the lock
//
// This is for emergency use only and should be used with caution
func (b *LayerCommandLock) Clear() error {
	cmd := commands.LayerCommandLock{
		LayerCommand: b.baseLayerChannelCommand(),
		Action:       types.LockActionClear,
		Passphrase:   nil,
	}
	return b.sendCommand(cmd)
}

type LayerCommandChannelInfo struct {
	LayerChannelBuilder
}

func (b *LayerChannelBuilder) Info() *LayerCommandChannelInfo {
	return &LayerCommandChannelInfo{
		LayerChannelBuilder: *b,
	}
}

// Generic gets information about the channel.
func (b *LayerCommandChannelInfo) Generic() (responses.QueryChannelInfoVerbose, error) {
	cmd := commands.LayerCommandInfo{
		LayerCommand: b.baseLayerChannelCommand(),
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	var infoChannel responses.QueryChannelInfoVerbose
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	return infoChannel, nil
}

// Delay get the current delay on the specified channel.
//
// Deprecated: This command does not return what it states as of CasparCG 2.5.0
// https://github.com/CasparCG/server/issues/1151
func (b *LayerCommandChannelInfo) Delay() (responses.QueryChannelInfoVerbose, error) {
	cmd := commands.LayerCommandInfoDelay{
		LayerCommand: b.baseLayerChannelCommand(),
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	fmt.Print(data)
	var infoChannel responses.QueryChannelInfoVerbose
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	return infoChannel, nil
}

//
// Layer Commands
//

type LayerLayerBuilder struct {
	LayerChannelBuilder

	layer int
}

func (c *LayerChannelBuilder) Layer(layer int) *LayerLayerBuilder {
	return &LayerLayerBuilder{
		LayerChannelBuilder: *c,
		layer:               layer,
	}
}

func (b *LayerLayerBuilder) baseLayerLayerCommand() commands.LayerCommand {
	return commands.LayerCommand{
		VideoChannel: b.videoChannel,
		Layer:        &b.layer,
	}
}

// Load loads a clip to the layer.
func (b *LayerLayerBuilder) Load(params types.LayerLoad) error {
	cmd := commands.LayerCommandLoad{
		LayerCommand: b.baseLayerLayerCommand(),
		Clip:         params.ClipName,
		Parameters:   params.Parameters,
	}
	return b.sendCommand(cmd)
}

// Play plays content on the layer.
func (b *LayerLayerBuilder) Play(params types.LayerPlay) error {
	cmd := commands.LayerCommandPlay{
		LayerCommand: b.baseLayerLayerCommand(),
		Clip:         params.ClipName,
		Parameters:   params.Parameters,
	}
	return b.sendCommand(cmd)
}

// Pause pauses playback on the layer.
func (b *LayerLayerBuilder) Pause() error {
	cmd := commands.LayerCommandPause{
		LayerCommand: b.baseLayerLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Resume resumes playback on the layer.
func (b *LayerLayerBuilder) Resume() error {
	cmd := commands.LayerCommandResume{
		LayerCommand: b.baseLayerLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Stop stops playback on the layer.
func (b *LayerLayerBuilder) Stop() error {
	cmd := commands.LayerCommandStop{
		LayerCommand: b.baseLayerLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Clear clears the layer.
func (b *LayerLayerBuilder) Clear() error {
	cmd := commands.LayerCommandClear{
		LayerCommand: b.baseLayerLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Call calls a function on the layer.
//
// TODO: Implement all possible parameter for CALL command
func (b *LayerLayerBuilder) Call(params []string) error {
	cmd := commands.LayerCommandCall{
		LayerCommand: b.baseLayerLayerCommand(),
		Params:       params,
	}
	return b.sendCommand(cmd)
}

// Swap swaps layers between channels.
func (b *LayerLayerBuilder) Swap(channel2 int, layer2 int, transforms bool) error {
	cmd := commands.LayerCommandSwap{
		LayerCommand:  b.baseLayerLayerCommand(),
		VideoChannel2: channel2,
		Layer2:        &layer2,
		Transform:     transforms,
	}
	return b.sendCommand(cmd)
}

type LayerCommandLayerInfo struct {
	LayerLayerBuilder
}

func (b *LayerLayerBuilder) Info() *LayerCommandLayerInfo {
	return &LayerCommandLayerInfo{
		LayerLayerBuilder: *b,
	}
}

// Generic gets information about the channel.
func (b *LayerCommandLayerInfo) Generic() (responses.QueryChannelInfoVerbose, error) {
	cmd := commands.LayerCommandInfo{
		LayerCommand: b.baseLayerLayerCommand(),
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	var infoChannel responses.QueryChannelInfoVerbose
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	return infoChannel, nil
}

// Delay get the current delay on the specified channel.
//
// Deprecated: This command does not return what it states as of CasparCG 2.5.0
// https://github.com/CasparCG/server/issues/1151
func (b *LayerCommandLayerInfo) Delay() (responses.QueryChannelInfoVerbose, error) {
	cmd := commands.LayerCommandInfoDelay{
		LayerCommand: b.baseLayerLayerCommand(),
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	fmt.Print(data)
	var infoChannel responses.QueryChannelInfoVerbose
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.QueryChannelInfoVerbose{}, err
	}
	return infoChannel, nil
}
