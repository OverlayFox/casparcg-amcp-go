package types

import (
	"fmt"
	"strconv"
)

type Fade struct {
	Duration int       // Duration defines the duration of the fade in frames. (e.g: 2sec * 50fps = 100 duration)
	Tween    TweenType // Tween defines the tween type to use for the fade. If not specified, it defaults to linear.
}

type MixerParamsFill struct {
	X      float32 // X defines the new x position, 0 = left edge of monitor, 0.5 = middle of monitor, 1.0 = right edge of monitor. Higher and lower values allowed.
	Y      float32 // Y defines The new y position, 0 = top edge of monitor, 0.5 = middle of monitor, 1.0 = bottom edge of monitor. Higher and lower values allowed.
	XScale float32 // XScale defines the new x scale, 1 = 1x the screen width, 0.5 = half the screen width. Higher and lower values allowed. Negative values flips the layer.
	YScale float32 // YScale defines the new y scale, 1 = 1x the screen height, 0.5 = half the screen height. Higher and lower values allowed. Negative values flips the layer.
}

type MixerLevels struct {
	MinInput float32 // MinInput defines the minimum input value (between 0 and 1) to accept RGB values within.
	MaxInput float32 // MaxInput defines the maximum input value (between 0 and 1) to accept RGB values within.

	Gamma float32 // Gamma adjusts the gamma of the image.

	MinOutput float32 // MinOutput defines the minimum output value (between 0 and 1) to output RGB values within.
	MaxOutput float32 // MaxOutput defines the maximum output value (between 0 and 1) to output RGB values within.
}

func MixerInfoLevelsFromResponse(data []string) (MixerLevels, error) {
	if len(data) < 5 {
		return MixerLevels{}, fmt.Errorf("unexpected response length: got %d, expected at least 5", len(data))
	}

	minInput, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerLevels{}, fmt.Errorf("invalid MinInput value: %w", err)
	}
	maxInput, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerLevels{}, fmt.Errorf("invalid MaxInput value: %w", err)
	}
	gamma, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerLevels{}, fmt.Errorf("invalid Gamma value: %w", err)
	}
	minOutput, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerLevels{}, fmt.Errorf("invalid MinOutput value: %w", err)
	}
	maxOutput, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerLevels{}, fmt.Errorf("invalid MaxOutput value: %w", err)
	}

	return MixerLevels{
		MinInput:  float32(minInput),
		MaxInput:  float32(maxInput),
		Gamma:     float32(gamma),
		MinOutput: float32(minOutput),
		MaxOutput: float32(maxOutput),
	}, nil
}

type MixerCrop struct {
	LeftEdge   float32 // LeftEdge defines a value between 0 and 1 defining how far into the layer to crop from the left edge.
	TopEdge    float32 // TopEdge defines a value between 0 and 1 defining how far into the layer to crop from the top edge.
	RightEdge  float32 // RightEdge defines a value between 0 and 1 defining how far into the layer to crop from the right edge.
	BottomEdge float32 // BottomEdge defines a value between 0 and 1 defining how far into the layer to crop from the bottom edge.
}

type MixerPerspective struct {
	TopLeftX float32 // TopLeftX defines the x coordinate of the top left corner.
	TopLeftY float32 // TopLeftY defines the y coordinate of the top left corner.

	TopRightX float32 // TopRightX defines the x coordinate of the top right corner.
	TopRightY float32 // TopRightY defines the y coordinate of the top right corner.

	BottomRightX float32 // BottomRightX defines the x coordinate of the bottom right corner.
	BottomRightY float32 // BottomRightY defines the y coordinate of the bottom right corner.

	BottomLeftX float32 // BottomLeftX defines the x coordinate of the bottom left corner.
	BottomLeftY float32 // BottomLeftY defines the y coordinate of the bottom left corner.
}

type CGAdd struct {
	Template   string // Template defines the name of the template to add. This is the filename of the template without the extension
	PlayOnLoad bool   // PlayOnLoad defines whether the template should start playing immediately when added.

	Data *string // Data defines optional data to pass to the template. This can be a JSON or XML inline string
}

type LayerLoad struct {
	ClipName string // ClipName defines the name of the clip to load. This is the filename of the clip without the extension.

	Parameters *[]string // Parameters defines optional parameters to pass to the clip. This can be used for dynamic templates or to pass other parameters to the clip.
}

type LayerPlay struct {
	ClipName *string // ClipName defines the name of the clip to play. This is the filename of the clip without the extension. If not specified, it will play the currently loaded clip.

	Parameters *[]string // Parameters defines optional parameters to pass to the clip. This can be used for dynamic templates or to pass other parameters to the clip.
}

type LayerAdd struct {
	ConsumerName string // ConsumerName defines the name of the consumer to add. This is the name of the consumer, e.g. "STREAM", "RECORD", etc. // TODO: Make this an enum of possible consumer types

	ConsumerIdx *int      // ConsumerIdx overrides the index that the consumer itself decides and can later be used with the REMOVE command to remove the consumer.
	Parameters  *[]string // Parameters are specific to the consumer being added. For example, for a STREAM consumer you can add []string{"udp://localhost:5004", "-vcodec", "libx264", "-tune", "zerolatency"}
}

type LayerRemove struct {
	ConsumerIdx *int      // ConsumerIdx overrides the index that the consumer itself decides and can later be used with the REMOVE command to remove the consumer.
	Parameters  *[]string // Parameters are specific to the consumer being added. For example, for a STREAM consumer you can add []string{"udp://localhost:5004", "-vcodec", "libx264", "-tune", "zerolatency"}
}

type LayerSet struct {
	VariableName string // VariableName defines the name of the variable to set.
	Value        string // Value defines the value to set the variable to.
}
