package casparcg

import (
	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

// CGBuilder provides a fluent interface for building CG (template) commands.
type CGBuilder struct {
	client *Client
}

// sendCommand abstracts sending a command that does not expect a response value.
func (b *CGBuilder) sendCommand(cmd interface{ String() string }) error {
	_, err := b.client.Send(cmd)
	return err
}

// CG creates a new CG command builder for the specified channel and layer.
func (c *Client) CG() *CGBuilder {
	return &CGBuilder{
		client: c,
	}
}

//
// Channel Commands
//

type CGChannelBuilder struct {
	CGBuilder

	videoChannel int
}

// Channel selects the video channel to operate on and returns a CGChannelBuilder for that channel.
func (c *CGBuilder) Channel(videoChannel int) *CGChannelBuilder {
	return &CGChannelBuilder{
		CGBuilder: CGBuilder{
			client: c.client,
		},
		videoChannel: videoChannel,
	}
}

func (b *CGChannelBuilder) baseCGChannelCommand() commands.CGCommand {
	return commands.CGCommand{
		VideoChannel: b.videoChannel,
		Layer:        nil, // Layer is nil for channel-level commands
		CgLayer:      nil, // CgLayer is nil for channel-level commands
	}
}

//
// Layer Commands
//

type CGLayerBuilder struct {
	CGChannelBuilder

	layer int
}

// Layer selects the layer to operate on and returns a CGLayerBuilder for that layer.
func (cb *CGChannelBuilder) Layer(layer int) *CGLayerBuilder {
	return &CGLayerBuilder{
		CGChannelBuilder: *cb,
		layer:            layer,
	}
}

func (b *CGLayerBuilder) baseCGLayerCommand() commands.CGCommand {
	cmd := b.baseCGChannelCommand()
	cmd.Layer = &b.layer
	return cmd
}

// Info retrieves information about the template on the specified layer.
//
// TODO: Implement response object of INFO command
func (b *CGLayerBuilder) Info(cgLayer *int) error {
	cmd := commands.TemplateCommandCGInvoke{
		CGCommand: b.baseCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}

//
// CG Layer Commands.
//

type CGCGLayerBuilder struct {
	CGLayerBuilder

	CgLayer int
}

// CGLayer selects the CG layer to operate on and returns a CGCGLayerBuilder for that CG layer.
func (cb *CGLayerBuilder) CGLayer(cgLayer int) *CGCGLayerBuilder {
	return &CGCGLayerBuilder{
		CGLayerBuilder: *cb,
		CgLayer:        cgLayer,
	}
}

func (b *CGCGLayerBuilder) baseCGCGLayerCommand() commands.CGCommand {
	cmd := b.baseCGLayerCommand()
	cmd.CgLayer = &b.CgLayer
	return cmd
}

// Add prepares a template for displaying.
func (b *CGCGLayerBuilder) Add(params types.CGAdd) error {
	cmd := commands.TemplateCommandCGAdd{
		CGCommand:  b.baseCGCGLayerCommand(),
		Template:   params.Template,
		PlayOnLoad: params.PlayOnLoad,
		Data:       params.Data,
	}
	return b.sendCommand(cmd)
}

// Play plays and displays the template in the specified layer.
func (b *CGCGLayerBuilder) Play() error {
	cmd := commands.TemplateCommandCGPlay{
		CGCommand: b.baseCGCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Stop stops the template in the specified layer.
func (b *CGCGLayerBuilder) Stop() error {
	cmd := commands.TemplateCommandCGStop{
		CGCommand: b.baseCGCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Next triggers a "continue" in the template.
func (b *CGCGLayerBuilder) Next() error {
	cmd := commands.TemplateCommandCGNext{
		CGCommand: b.baseCGCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Remove removes the template from the specified layer.
func (b *CGCGLayerBuilder) Remove() error {
	cmd := commands.TemplateCommandCGRemove{
		CGCommand: b.baseCGCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Clear removes all templates on the video layer.
func (b *CGCGLayerBuilder) Clear() error {
	cmd := commands.TemplateCommandCGClear{
		CGCommand: b.baseCGCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}

// Update sends new data to the template on specified layer.
//
// data - string: data to pass to the template. This can be a JSON or XML inline string.
func (b *CGCGLayerBuilder) Update(data string) error {
	cmd := commands.TemplateCommandCGUpdate{
		CGCommand: b.baseCGCGLayerCommand(),
		Data:      data,
	}
	return b.sendCommand(cmd)
}

// Invoke invokes the given method on the template.
//
// method - string: the name of the method to invoke on the template.
func (b *CGCGLayerBuilder) Invoke(method string) error {
	cmd := commands.TemplateCommandCGInvoke{
		CGCommand: b.baseCGCGLayerCommand(),
		Method:    method,
	}
	return b.sendCommand(cmd)
}

// Info retrieves information about the template on the specified layer.
//
// TODO: Implement response object of INFO command
func (b *CGCGLayerBuilder) Info(cgLayer *int) error {
	cmd := commands.TemplateCommandCGInvoke{
		CGCommand: b.baseCGCGLayerCommand(),
	}
	return b.sendCommand(cmd)
}
