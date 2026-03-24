package types

import (
	"fmt"
	"strconv"
)

type MixerCommand struct {
	VideoChannel int
	Layer        int // defaults to 9999
}

type Fade struct {
	Duration int // in frames
	Tween    TweenType
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
	Tween        *TweenType
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
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}

	return cmd
}

type MixerCommandBlend struct {
	MixerCommand

	BlendMode *BlendMode
}

func (c MixerCommandBlend) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d BLEND", c.VideoChannel, c.Layer)
	if c.BlendMode != nil {
		cmd += " " + c.BlendMode.String()
	}
	return cmd
}

type MixerCommandInvert struct {
	MixerCommand

	Invert *bool
}

func (c MixerCommandInvert) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d INVERT", c.VideoChannel, c.Layer)
	if c.Invert != nil {
		if *c.Invert {
			cmd += " 1"
		} else {
			cmd += " 0"
		}
	}
	return cmd
}

type MixerCommandOpacity struct {
	MixerCommand

	Opacity *float32

	Duration *int
	Tween    *TweenType
}

func (c MixerCommandOpacity) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d OPACITY", c.VideoChannel, c.Layer)
	if c.Opacity != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Opacity)
	}
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}

type MixerCommandBrightness struct {
	MixerCommand

	Brightness *float32

	Duration *int
	Tween    *TweenType
}

func (c MixerCommandBrightness) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d BRIGHTNESS", c.VideoChannel, c.Layer)
	if c.Brightness != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Brightness)
	}
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}

type MixerCommandSaturation struct {
	MixerCommand

	Saturation *float32

	Duration *int
	Tween    *TweenType
}

func (c MixerCommandSaturation) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d SATURATION", c.VideoChannel, c.Layer)
	if c.Saturation != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Saturation)
	}
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}
