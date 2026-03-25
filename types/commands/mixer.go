package commands

import (
	"fmt"
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
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
	Tween        *types.TweenType
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

	BlendMode *types.BlendMode
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
	Tween    *types.TweenType
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
	Tween    *types.TweenType
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
	Tween    *types.TweenType
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

type MixerCommandContrast struct {
	MixerCommand

	Contrast *float32

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandContrast) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d CONTRAST", c.VideoChannel, c.Layer)
	if c.Contrast != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Contrast)
	}
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}

type MixerCommandLevels struct {
	MixerCommand

	MinInput  *float32 // MinInput and MaxInput define the input range (between 0 and 1) to accept RGB values within.
	MaxInput  *float32
	Gamma     *float32 // Gamma adjusts the gamma of the image.
	MinOutput *float32 // MinOutput and MaxOutput define the output range (between 0 and 1) to scale the accepted input RGB values to.
	MaxOutput *float32

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandLevels) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d LEVELS", c.VideoChannel, c.Layer)
	if c.MinInput != nil {
		cmd += " " + fmt.Sprintf("%f", *c.MinInput)
	}
	if c.MaxInput != nil {
		cmd += " " + fmt.Sprintf("%f", *c.MaxInput)
	}
	if c.Gamma != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Gamma)
	}
	if c.MinOutput != nil {
		cmd += " " + fmt.Sprintf("%f", *c.MinOutput)
	}
	if c.MaxOutput != nil {
		cmd += " " + fmt.Sprintf("%f", *c.MaxOutput)
	}
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}

type MixerCommandFill struct {
	MixerCommand

	X      *float32 // X the new x position, 0 = left edge of monitor, 0.5 = middle of monitor, 1.0 = right edge of monitor. Higher and lower values allowed.
	Y      *float32 // Y the new y position, 0 = top edge of monitor, 0.5 = middle of monitor, 1.0 = bottom edge of monitor. Higher and lower values allowed.
	XScale *float32 // XScale the new x scale, 1.0 = original size, 0.5 = half size, 2.0 = double size. Higher and lower values allowed.
	YScale *float32 // YScale the new y scale, 1.0 = original size, 0.5 = half size, 2.0 = double size. Higher and lower values allowed.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandFill) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d FILL", c.VideoChannel, c.Layer)
	if c.X != nil {
		cmd += " " + fmt.Sprintf("%f", *c.X)
	}
	if c.Y != nil {
		cmd += " " + fmt.Sprintf("%f", *c.Y)
	}
	if c.XScale != nil {
		cmd += " " + fmt.Sprintf("%f", *c.XScale)
	}
	if c.YScale != nil {
		cmd += " " + fmt.Sprintf("%f", *c.YScale)
	}
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}

type MixerClip struct {
	MixerCommand

	X      float32 // X: The new x position, 0 = left edge of monitor, 0.5 = middle of monitor, 1.0 = right edge of monitor. Higher and lower values allowed.
	Y      float32 // Y: The new y position, 0 = top edge of monitor, 0.5 = middle of monitor, 1.0 = bottom edge of monitor. Higher and lower values allowed.
	Width  float32 // Width: The new width, 1 = 1x the screen width, 0.5 = half the screen width. Higher and lower values allowed. Negative values flips the layer.
	Height float32 // Height: The new height, 1 = 1x the screen height, 0.5 = half the screen height. Higher and lower values allowed. Negative values flips the layer.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerClip) String() string {
	cmd := fmt.Sprintf("MIXER %d-%d CLIP", c.VideoChannel, c.Layer)
	cmd += " " + fmt.Sprintf("%f", c.X)
	cmd += " " + fmt.Sprintf("%f", c.Y)
	cmd += " " + fmt.Sprintf("%f", c.Width)
	cmd += " " + fmt.Sprintf("%f", c.Height)
	if c.Duration != nil {
		cmd += " " + strconv.Itoa(*c.Duration)
	}
	if c.Tween != nil {
		cmd += " " + c.Tween.String()
	}
	return cmd
}
