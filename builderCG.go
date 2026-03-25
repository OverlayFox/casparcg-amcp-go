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

// Add prepares a template for displaying.
func (b *CGBuilder) Add(cgLayer int, template string, playOnLoad bool, data *string) error {
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

// Play plays and displays the template in the specified layer.
func (b *CGBuilder) Play(cgLayer int) error {
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

// Stop stops the template in the specified layer.
func (b *CGBuilder) Stop(cgLayer int) error {
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

// Next triggers a "continue" in the template.
func (b *CGBuilder) Next(cgLayer int) error {
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

// Remove removes the template from the specified layer.
func (b *CGBuilder) Remove(cgLayer int) error {
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

// Clear removes all templates on the video layer.
func (b *CGBuilder) Clear() error {
	cmd := commands.TemplateCommandCGClear{
		TemplateCommandCG: commands.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// Update sends new data to the template on specified layer.
func (b *CGBuilder) Update(cgLayer int, data string) error {
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

// Invoke invokes the given method on the template.
func (b *CGBuilder) Invoke(cgLayer int, method string) error {
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

// Info retrieves information about the template on the specified layer.
func (b *CGBuilder) Info(cgLayer *int) error {
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
