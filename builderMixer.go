package casparcg

import (
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
	"github.com/overlayfox/casparcg-amcp-go/types/responses"
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
	cmd := commands.MixerCommandKeyer{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return false, err
	}

	return responses.BoolFromResponse(resp)
}

// Keyer replaces layer n+1's alpha with the R (red) channel of layer n, and hides the RGB channels of layer n.
// If show is true then the specified layer will not be rendered, instead it will be used as the key for the layer above.
func (b *MixerBuilder) Keyer(show bool) error {
	cmd := commands.MixerCommandKeyer{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Show: &show,
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetChromaInfo() (responses.MixerInfoChroma, error) {
	cmd := commands.MixerCommandChroma{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerInfoChroma{}, err
	}

	return responses.MixerInfoChromaFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

func (b *MixerBuilder) ChromaEnable(params responses.MixerInfoChroma, fade *types.Fade) error {
	enable := true
	cmd := commands.MixerCommandChroma{
		MixerCommand: commands.MixerCommand{
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
	cmd := commands.MixerCommandChroma{
		MixerCommand: commands.MixerCommand{
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
	cmd := commands.MixerCommandBlend{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return "", err
	}

	return responses.MixerBlendModeFromResponse(resp)
}

func (b *MixerBuilder) SetBlendMode(mode types.BlendMode) error {
	cmd := commands.MixerCommandBlend{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		BlendMode: &mode,
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetInvertState() (bool, error) {
	cmd := commands.MixerCommandInvert{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return false, err
	}

	return responses.BoolFromResponse(resp)
}

func (b *MixerBuilder) SetInvert(state bool) error {
	cmd := commands.MixerCommandInvert{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Invert: &state,
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetOpacity() (float32, error) {
	cmd := commands.MixerCommandOpacity{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}

	return responses.FloatFromResponse(resp)
}

func (b *MixerBuilder) SetOpacity(opacity float32, fade *types.Fade) error {
	cmd := commands.MixerCommandOpacity{
		MixerCommand: commands.MixerCommand{
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
	cmd := commands.MixerCommandBrightness{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}

	return responses.FloatFromResponse(resp)
}

func (b *MixerBuilder) SetBrightness(brightness float32, fade *types.Fade) error {
	cmd := commands.MixerCommandBrightness{
		MixerCommand: commands.MixerCommand{
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

func (b *MixerBuilder) GetSaturation() (float32, error) {
	cmd := commands.MixerCommandSaturation{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}

	return responses.FloatFromResponse(resp)
}

func (b *MixerBuilder) SetSaturation(saturation float32, fade *types.Fade) error {
	cmd := commands.MixerCommandSaturation{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Saturation: &saturation,
	}
	if fade != nil {
		cmd.Duration = &fade.Duration
		cmd.Tween = &fade.Tween
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetContrast() (float32, error) {
	cmd := commands.MixerCommandContrast{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}

	return responses.FloatFromResponse(resp)
}

func (b *MixerBuilder) SetContrast(contrast float32, fade *types.Fade) error {
	cmd := commands.MixerCommandContrast{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Contrast: &contrast,
	}
	if fade != nil {
		cmd.Duration = &fade.Duration
		cmd.Tween = &fade.Tween
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetLevels() (types.MixerInfoLevels, error) {
	cmd := commands.MixerCommandLevels{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return types.MixerInfoLevels{}, err
	}

	return types.MixerInfoLevelsFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

func (b *MixerBuilder) SetLevels(params types.MixerInfoLevels, fade *types.Fade) error {
	cmd := commands.MixerCommandLevels{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		MinInput:  &params.MinInput,
		MaxInput:  &params.MaxInput,
		Gamma:     &params.Gamma,
		MinOutput: &params.MinOutput,
		MaxOutput: &params.MaxOutput,
	}
	if fade != nil {
		cmd.Duration = &fade.Duration
		cmd.Tween = &fade.Tween
	}
	_, err := b.client.Send(cmd)
	return err
}

func (b *MixerBuilder) GetFill() (responses.MixerInfoFill, error) {
	cmd := commands.MixerCommandFill{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerInfoFill{}, err
	}

	return responses.MixerInfoFillFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetFill scales/positions the video stream on the specified layer.
// The positioning and scaling is done around the anchor point set by MIXER ANCHOR.
func (b *MixerBuilder) SetFill(params types.MixerParamsFill, fade *types.Fade) error {
	cmd := commands.MixerCommandFill{
		MixerCommand: commands.MixerCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		X:      params.X,
		Y:      params.Y,
		XScale: params.XScale,
		YScale: params.YScale,
	}
	if fade != nil {
		cmd.Duration = &fade.Duration
		cmd.Tween = &fade.Tween
	}
	_, err := b.client.Send(cmd)
	return err
}
