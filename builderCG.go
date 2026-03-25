package casparcg

import (
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

// CGBuilder provides a fluent interface for building CG (template) commands.
type CGBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

// CG creates a new CG command builder for the specified channel and layer
// Example: client.CG(1, 12).STOP(2).
func (c *Client) CG(videoChannel, layer int) *CGBuilder {
	return &CGBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// ADD prepares a template for displaying.
func (b *CGBuilder) ADD(cgLayer int, template string, playOnLoad bool, data *string) error {
	cmd := commands.TemplateCommandCGAdd{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer:    cgLayer,
		Template:   template,
		PlayOnLoad: playOnLoad,
		Data:       data,
	}
	_, err := b.client.Send(cmd)
	return err
}

// PLAY plays and displays the template in the specified layer.
func (b *CGBuilder) PLAY(cgLayer int) error {
	cmd := commands.TemplateCommandCGPlay{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// STOP stops the template in the specified layer.
func (b *CGBuilder) STOP(cgLayer int) error {
	cmd := commands.TemplateCommandCGStop{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// NEXT triggers a "continue" in the template.
func (b *CGBuilder) NEXT(cgLayer int) error {
	cmd := commands.TemplateCommandCGNext{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// REMOVE removes the template from the specified layer.
func (b *CGBuilder) REMOVE(cgLayer int) error {
	cmd := commands.TemplateCommandCGRemove{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// CLEAR removes all templates on the video layer.
func (b *CGBuilder) CLEAR() error {
	cmd := commands.TemplateCommandCGClear{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// UPDATE sends new data to the template on specified layer.
func (b *CGBuilder) UPDATE(cgLayer int, data string) error {
	cmd := commands.TemplateCommandCGUpdate{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
		Data:    data,
	}
	_, err := b.client.Send(cmd)
	return err
}

// INVOKE invokes the given method on the template.
func (b *CGBuilder) INVOKE(cgLayer int, method string) error {
	cmd := commands.TemplateCommandCGInvoke{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
		Method:  method,
	}
	_, err := b.client.Send(cmd)
	return err
}

// INFO retrieves information about the template on the specified layer.
func (b *CGBuilder) INFO(cgLayer *int) error {
	cmd := commands.TemplateCommandCGInfo{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}
