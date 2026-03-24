package casparcg

import (
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/returns"
)

type MixerBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

func (c *Client) Mixer(videoChannel, layer int) *MixerBuilder {
	return &MixerBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

func (b *MixerBuilder) GetKeyerState() (bool, error) {
	cmd := types.MixerCommandKeyer{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return false, err
	}

	return returns.BoolFromResponse(resp)
}

// Keyer replaces layer n+1's alpha with the R (red) channel of layer n, and hides the RGB channels of layer n.
// If show is true then the specified layer will not be rendered, instead it will be used as the key for the layer above.
func (b *MixerBuilder) Keyer(show bool) error {
	cmd := types.MixerCommandKeyer{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Show: &show,
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetChromaInfo() (returns.MixerChromaInfo, error) {
	cmd := types.MixerCommandChroma{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return returns.MixerChromaInfo{}, err
	}

	return returns.MixerChromaInfoFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

func (b *MixerBuilder) ChromaEnable(params returns.MixerChromaInfo, fade *types.Fade) error {
	enable := true
	cmd := types.MixerCommandChroma{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Enable:                  &enable,
		TargetHue:               &params.TargetHue,
		HueWidth:                &params.HueWidth,
		MinSaturation:           &params.MinSaturation,
		MinBrightness:           &params.MinBrightness,
		Softness:                &params.Softness,
		SpillSuppress:           &params.SpillSuppress,
		SpillSuppressSaturation: &params.SpillSuppressSaturation,
		ShowMask:                &params.ShowMask,
	}

	if fade != nil {
		cmd.FadeDuration = &fade.Duration
		cmd.Tween = &fade.Tween
	}

	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) ChromaDisable(fade *types.Fade) error {
	enable := false
	cmd := types.MixerCommandChroma{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Enable: &enable,
	}

	if fade != nil {
		cmd.FadeDuration = &fade.Duration
		cmd.Tween = &fade.Tween
	}

	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetBlendMode() (types.BlendMode, error) {
	cmd := types.MixerCommandBlend{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return "", err
	}

	return returns.MixerBlendModeFromResponse(resp)
}

func (b *MixerBuilder) SetBlendMode(mode types.BlendMode) error {
	cmd := types.MixerCommandBlend{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		BlendMode: &mode,
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetInvertState() (bool, error) {
	cmd := types.MixerCommandInvert{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return false, err
	}

	return returns.BoolFromResponse(resp)
}

func (b *MixerBuilder) SetInvert(state bool) error {
	cmd := types.MixerCommandInvert{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Invert: &state,
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetOpacity() (float32, error) {
	cmd := types.MixerCommandOpacity{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}

	return returns.FloatFromResponse(resp)
}

func (b *MixerBuilder) SetOpacity(opacity float32, fade *types.Fade) error {
	cmd := types.MixerCommandOpacity{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Opacity: &opacity,
	}
	if fade != nil {
		cmd.Duration = &fade.Duration
		cmd.Tween = &fade.Tween
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetBrightness() (float32, error) {
	cmd := types.MixerCommandBrightness{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}

	return returns.FloatFromResponse(resp)
}

func (b *MixerBuilder) SetBrightness(brightness float32, fade *types.Fade) error {
	cmd := types.MixerCommandBrightness{
		MixerCommand: types.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Brightness: &brightness,
	}
	if fade != nil {
		cmd.Duration = &fade.Duration
		cmd.Tween = &fade.Tween
	}
	_, err := b.client.Send(cmd)
	return err
}
