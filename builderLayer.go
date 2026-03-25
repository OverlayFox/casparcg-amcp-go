package casparcg

import (
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

// LayerBuilder provides a fluent interface for building layer-based commands.
type LayerBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

// Layer creates a new layer command builder for the specified channel and layer
// Example: client.Layer(1, 10).PLAY("myclip", nil).
func (c *Client) Layer(videoChannel, layer int) *LayerBuilder {
	return &LayerBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// Load loads a clip to the layer.
func (b *LayerBuilder) Load(clip string, parameters *map[string]string) error {
	cmd := commands.CommandLoad{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Clip:       clip,
		Parameters: parameters,
	}
	_, err := b.client.Send(cmd)
	return err
}

// Play plays content on the layer.
func (b *LayerBuilder) Play(clip *string, parameters *map[string]string) error {
	cmd := commands.CommandPlay{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Clip:       clip,
		Parameters: parameters,
	}
	_, err := b.client.Send(cmd)
	return err
}

// Pause pauses playback on the layer.
func (b *LayerBuilder) Pause() error {
	cmd := commands.CommandPause{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// Resume resumes playback on the layer.
func (b *LayerBuilder) Resume() error {
	cmd := commands.CommandResume{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// Stop stops playback on the layer.
func (b *LayerBuilder) Stop() error {
	cmd := commands.CommandStop{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// Clear clears the layer.
func (b *LayerBuilder) Clear() error {
	cmd := commands.CommandClear{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// Call calls a function on the layer.
func (b *LayerBuilder) Call(params map[string]string) error {
	cmd := commands.CommandCall{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Params: params,
	}
	_, err := b.client.Send(cmd)
	return err
}
