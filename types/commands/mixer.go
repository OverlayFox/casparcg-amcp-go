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
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "KEYER")
	return appendBool(cmd, c.Show)
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
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "CHROMA")
	cmd = appendBool(cmd, c.Enable)
	cmd = appendFloat(cmd, c.TargetHue)
	cmd = appendFloat(cmd, c.HueWidth)
	cmd = appendFloat(cmd, c.MinSaturation)
	cmd = appendFloat(cmd, c.MinBrightness)
	cmd = appendFloat(cmd, c.Softness)
	cmd = appendFloat(cmd, c.SpillSuppress)
	cmd = appendFloat(cmd, c.SpillSuppressSaturation)
	cmd = appendBool(cmd, c.ShowMask)
	return appendDurationTween(cmd, c.FadeDuration, c.Tween)
}

type MixerCommandBlend struct {
	MixerCommand

	BlendMode *types.BlendMode
}

func (c MixerCommandBlend) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "BLEND")
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
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "INVERT")
	return appendBool(cmd, c.Invert)
}

type MixerCommandOpacity struct {
	MixerCommand

	Opacity *float32

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandOpacity) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "OPACITY")
	cmd = appendFloat(cmd, c.Opacity)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerCommandBrightness struct {
	MixerCommand

	Brightness *float32

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandBrightness) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "BRIGHTNESS")
	cmd = appendFloat(cmd, c.Brightness)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerCommandSaturation struct {
	MixerCommand

	Saturation *float32

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandSaturation) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "SATURATION")
	cmd = appendFloat(cmd, c.Saturation)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerCommandContrast struct {
	MixerCommand

	Contrast *float32

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCommandContrast) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "CONTRAST")
	cmd = appendFloat(cmd, c.Contrast)
	return appendDurationTween(cmd, c.Duration, c.Tween)
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
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "LEVELS")
	cmd = appendFloat(cmd, c.MinInput)
	cmd = appendFloat(cmd, c.MaxInput)
	cmd = appendFloat(cmd, c.Gamma)
	cmd = appendFloat(cmd, c.MinOutput)
	cmd = appendFloat(cmd, c.MaxOutput)
	return appendDurationTween(cmd, c.Duration, c.Tween)
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
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "FILL")
	cmd = appendFloat(cmd, c.X)
	cmd = appendFloat(cmd, c.Y)
	cmd = appendFloat(cmd, c.XScale)
	cmd = appendFloat(cmd, c.YScale)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerClip struct {
	MixerCommand

	X      *float32 // X: The new x position, 0 = left edge of monitor, 0.5 = middle of monitor, 1.0 = right edge of monitor. Higher and lower values allowed.
	Y      *float32 // Y: The new y position, 0 = top edge of monitor, 0.5 = middle of monitor, 1.0 = bottom edge of monitor. Higher and lower values allowed.
	Width  *float32 // Width: The new width, 1 = 1x the screen width, 0.5 = half the screen width. Higher and lower values allowed. Negative values flips the layer.
	Height *float32 // Height: The new height, 1 = 1x the screen height, 0.5 = half the screen height. Higher and lower values allowed. Negative values flips the layer.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerClip) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "CLIP")
	cmd = appendFloat(cmd, c.X)
	cmd = appendFloat(cmd, c.Y)
	cmd = appendFloat(cmd, c.Width)
	cmd = appendFloat(cmd, c.Height)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerAnchor struct {
	MixerCommand

	X *float32 // X: The x anchor point, 0 = left edge of layer, 0.5 = middle of layer, 1.0 = right edge of layer. Higher and lower values allowed.
	Y *float32 // Y: The y anchor point, 0 = top edge of layer, 0.5 = middle of layer, 1.0 = bottom edge of layer. Higher and lower values allowed.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerAnchor) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "ANCHOR")
	cmd = appendFloat(cmd, c.X)
	cmd = appendFloat(cmd, c.Y)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerCrop struct {
	MixerCommand

	LeftEdge   *float32 // LeftEdge: A value between 0 and 1 defining how far into the layer to crop from the left edge.
	TopEdge    *float32 // TopEdge: A value between 0 and 1 defining how far into the layer to crop from the top edge.
	RightEdge  *float32 // RightEdge: A value between 0 and 1 defining how far into the layer to crop from the right edge.
	BottomEdge *float32 // BottomEdge: A value between 0 and 1 defining how far into the layer to crop from the bottom edge.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerCrop) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "CROP")
	cmd = appendFloat(cmd, c.LeftEdge)
	cmd = appendFloat(cmd, c.TopEdge)
	cmd = appendFloat(cmd, c.RightEdge)
	cmd = appendFloat(cmd, c.BottomEdge)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerRotation struct {
	MixerCommand

	Angle *float32 // Angle: The absolute rotation angle in degrees going from 0 to 360.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerRotation) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "ROTATION")
	cmd = appendFloat(cmd, c.Angle)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerPerspective struct {
	MixerCommand

	TopLeftX *float32 // TopLeftX: Defines the x coordinate of the top left corner.
	TopLeftY *float32 // TopLeftY: Defines the y coordinate of the top left corner.

	TopRightX *float32 // TopRightX: Defines the x coordinate of the top right corner.
	TopRightY *float32 // TopRightY: Defines the y coordinate of the top right corner.

	BottomLeftX *float32 // BottomLeftX: Defines the x coordinate of the bottom left corner.
	BottomLeftY *float32 // BottomLeftY: Defines the y coordinate of the bottom left corner.

	BottomRightX *float32 // BottomRightX: Defines the x coordinate of the bottom right corner.
	BottomRightY *float32 // BottomRightY: Defines the y coordinate of the bottom right corner.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerPerspective) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "PERSPECTIVE")
	cmd = appendFloat(cmd, c.TopLeftX)
	cmd = appendFloat(cmd, c.TopLeftY)
	cmd = appendFloat(cmd, c.TopRightX)
	cmd = appendFloat(cmd, c.TopRightY)
	cmd = appendFloat(cmd, c.BottomLeftX)
	cmd = appendFloat(cmd, c.BottomLeftY)
	cmd = appendFloat(cmd, c.BottomRightX)
	cmd = appendFloat(cmd, c.BottomRightY)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

type MixerMipMap struct {
	MixerCommand

	Enable *bool
}

func (c MixerMipMap) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "MIPMAP")
	return appendBool(cmd, c.Enable)
}

type MixerVolume struct {
	MixerCommand

	Volume *float32 // Volume: The new volume, 1.0 = original volume, 0.5 = half volume, 2.0 = double volume. Higher and lower values allowed.

	Duration *int
	Tween    *types.TweenType
}

func (c MixerVolume) String() string {
	cmd := baseMixerCmd(c.VideoChannel, c.Layer, "VOLUME")
	cmd = appendFloat(cmd, c.Volume)
	return appendDurationTween(cmd, c.Duration, c.Tween)
}

//
// Helper functions
//

func baseMixerCmd(channel, layer int, name string) string {
	return fmt.Sprintf("MIXER %d-%d %s", channel, layer, name)
}

func appendFloat(cmd string, value *float32) string {
	if value != nil {
		return cmd + " " + fmt.Sprintf("%f", *value)
	}
	return cmd
}

func appendBool(cmd string, value *bool) string {
	if value != nil {
		if *value {
			return cmd + " 1"
		}
		return cmd + " 0"
	}
	return cmd
}

func appendDurationTween(cmd string, duration *int, tween *types.TweenType) string {
	if duration != nil {
		cmd += " " + strconv.Itoa(*duration)
	}
	if tween != nil {
		cmd += " " + tween.String()
	}
	return cmd
}
