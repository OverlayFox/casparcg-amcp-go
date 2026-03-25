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

// LOAD loads a clip to the layer.
func (b *LayerBuilder) LOAD(clip string, parameters *map[string]string) error {
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

// PLAY plays content on the layer.
func (b *LayerBuilder) PLAY(clip *string, parameters *map[string]string) error {
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

// PAUSE pauses playback on the layer.
func (b *LayerBuilder) PAUSE() error {
	cmd := commands.CommandPause{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// RESUME resumes playback on the layer.
func (b *LayerBuilder) RESUME() error {
	cmd := commands.CommandResume{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// STOP stops playback on the layer.
func (b *LayerBuilder) STOP() error {
	cmd := commands.CommandStop{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// CLEAR clears the layer.
func (b *LayerBuilder) CLEAR() error {
	cmd := commands.CommandClear{
		BasicCommand: commands.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// CALL calls a function on the layer.
func (b *LayerBuilder) CALL(params map[string]string) error {
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
