package types

import (
	"fmt"
	"strconv"
)

type MixerCommand struct {
	VideoChannel int
	Layer        int // defaults to 9999
}

type MixerCommandKeyer struct {
	MixerCommand

	Show *bool
}

func (c MixerCommandKeyer) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d KEYER", c.VideoChannel, c.Layer)
	if c.Show != nil {
		if *c.Show {
			cmd += " 1"
		} else {
			cmd += " 0"
		}
	}
	return cmd
}

type MixerCommandChroma struct {
	MixerCommand

	Enable *bool

	TargetHue               *float32
	HueWidth                *float32
	MinSaturation           *float32
	MinBrightness           *float32
	Softness                *float32
	SpillSuppress           *float32
	SpillSuppressSaturation *float32
	ShowMask                *bool

	FadeDuration *int
	FadeTween    *string
}

func (c MixerCommandChroma) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d CHROMA", c.VideoChannel, c.Layer)
	if c.Enable != nil {
		if *c.Enable {
			cmd += " 1"
		} else {
			cmd += " 0"
		}
	}

	if c.TargetHue != nil {
		cmd += " " + fmt.Sprintf("%f", *c.TargetHue)
	}
	if c.HueWidth != nil {
		cmd += " " + fmt.Sprintf("%f", *c.HueWidth)
	}
	if c.MinSaturation != nil {
		cmd += " " + fmt.Sprintf("%f", *c.MinSaturation)
	}
	if c.MinBrightness != nil {
		cmd += " " + fmt.Sprintf("%f", *c.MinBrightness)
	}
	if c.Softness != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Softness)
	}
	if c.SpillSuppress != nil {
		cmd += " " + fmt.Sprintf("%f", *c.SpillSuppress)
	}
	if c.SpillSuppressSaturation != nil {
		cmd += " " + fmt.Sprintf("%f", *c.SpillSuppressSaturation)
	}
	if c.ShowMask != nil {
		if *c.ShowMask {
			cmd += " 1"
		} else {
			cmd += " 0"
		}
	}

	if c.FadeDuration != nil {
		cmd += " " + strconv.Itoa(*c.FadeDuration)
	}
	if c.FadeTween != nil {
		cmd += " " + *c.FadeTween
	}

	return cmd
}

type BlendMode string

const (
	BlendModeNormal BlendMode = "normal"
	BlendModeScreen BlendMode = "screen"
)

func ParseBlendMode(s string) (BlendMode, error) {
	validBlendModes := map[BlendMode]any{
		BlendModeNormal: nil,
		BlendModeScreen: nil,
	}

	mode := BlendMode(s)
	if _, ok := validBlendModes[mode]; !ok {
		return "", fmt.Errorf("invalid blend mode: %s", s)
	}
	return mode, nil
}

type MixerCommandBlend struct {
	MixerCommand

	BlendMode *BlendMode
}

func (c MixerCommandBlend) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d BLEND", c.VideoChannel, c.Layer)
	if c.BlendMode != nil {
		cmd += " " + string(*c.BlendMode)
	}
	return cmd
}
