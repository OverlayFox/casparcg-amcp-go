package responses

import (
	"fmt"
)

type MixerChroma struct {
	Enabled                 bool
	TargetHue               float32
	HueWidth                float32
	MinSaturation           float32
	MinBrightness           float32
	Softness                float32
	SpillSuppress           float32
	SpillSuppressSaturation float32
	ShowMask                bool
}

func MixerChromaFromResponse(data []string) (MixerChroma, error) {
	fields := []string{
		"TargetHue",
		"HueWidth",
		"MinSaturation",
		"MinBrightness",
		"Softness",
		"SpillSuppress",
		"SpillSuppressSaturation",
	}

	if len(data) < len(fields)+2 { // +2 for enabled and showMask
		return MixerChroma{}, fmt.Errorf("unexpected response length: got %d, expected at least %d", len(data), len(fields)+2)
	}

	values, err := parseFloat32Slice(data[1:8], fields)
	if err != nil {
		return MixerChroma{}, err
	}

	return MixerChroma{
		Enabled:                 parseBool(data[0]),
		TargetHue:               values[0],
		HueWidth:                values[1],
		MinSaturation:           values[2],
		MinBrightness:           values[3],
		Softness:                values[4],
		SpillSuppress:           values[5],
		SpillSuppressSaturation: values[6],
		ShowMask:                parseBool(data[8]),
	}, nil
}

type MixerFill struct {
	X      float32
	Y      float32
	XScale float32
	YScale float32
}

func MixerFillFromResponse(data []string) (MixerFill, error) {
	values, err := parseFloat32Slice(data, []string{"X", "Y", "XScale", "YScale"})
	if err != nil {
		return MixerFill{}, err
	}

	return MixerFill{
		X:      values[0],
		Y:      values[1],
		XScale: values[2],
		YScale: values[3],
	}, nil
}

type MixerClip struct {
	X      float32 // X: The new x position, 0 = left edge of monitor, 0.5 = middle of monitor, 1.0 = right edge of monitor. Higher and lower values allowed.
	Y      float32 // Y: The new y position, 0 = top edge of monitor, 0.5 = middle of monitor, 1.0 = bottom edge of monitor. Higher and lower values allowed.
	Width  float32 // Width: The new width, 1 = 1x the screen width, 0.5 = half the screen width. Higher and lower values allowed. Negative values flips the layer.
	Height float32 // Height: The new height, 1 = 1x the screen height, 0.5 = half the screen height. Higher and lower values allowed. Negative values flips the layer.
}

func MixerClipFromResponse(data []string) (MixerClip, error) {
	values, err := parseFloat32Slice(data, []string{"X", "Y", "Width", "Height"})
	if err != nil {
		return MixerClip{}, err
	}

	return MixerClip{
		X:      values[0],
		Y:      values[1],
		Width:  values[2],
		Height: values[3],
	}, nil
}

type MixerAnchor struct {
	X float32 // X defines the x anchor point, 0 = left edge of layer, 0.5 = middle of layer, 1.0 = right edge of layer. Higher and lower values allowed.
	Y float32 // Y defines the y anchor point, 0 = top edge of layer, 0.5 = middle of layer, 1.0 = bottom edge of layer. Higher and lower values allowed.
}

func MixerAnchorFromResponse(data []string) (MixerAnchor, error) {
	values, err := parseFloat32Slice(data, []string{"X", "Y"})
	if err != nil {
		return MixerAnchor{}, err
	}

	return MixerAnchor{
		X: values[0],
		Y: values[1],
	}, nil
}

type MixerCrop struct {
	LeftEdge   float32
	TopEdge    float32
	RightEdge  float32
	BottomEdge float32
}

func MixerCropFromResponse(data []string) (MixerCrop, error) {
	values, err := parseFloat32Slice(data, []string{"LeftEdge", "TopEdge", "RightEdge", "BottomEdge"})
	if err != nil {
		return MixerCrop{}, err
	}

	return MixerCrop{
		LeftEdge:   values[0],
		TopEdge:    values[1],
		RightEdge:  values[2],
		BottomEdge: values[3],
	}, nil
}

type MixerPerspective struct {
	TopLeftX float32 // TopLeftX: Defines the x coordinate of the top left corner.
	TopLeftY float32 // TopLeftY: Defines the y coordinate of the top left corner.

	TopRightX float32 // TopRightX: Defines the x coordinate of the top right corner.
	TopRightY float32 // TopRightY: Defines the y coordinate of the top right corner.

	BottomLeftX float32 // BottomLeftX: Defines the x coordinate of the bottom left corner.
	BottomLeftY float32 // BottomLeftY: Defines the y coordinate of the bottom left corner.

	BottomRightX float32 // BottomRightX: Defines the x coordinate of the bottom right corner.
	BottomRightY float32 // BottomRightY: Defines the y coordinate of the bottom right corner.
}

func MixerPerspectiveFromResponse(data []string) (MixerPerspective, error) {
	fields := []string{"TopLeftX", "TopLeftY", "TopRightX", "TopRightY", "BottomLeftX", "BottomLeftY", "BottomRightX", "BottomRightY"}
	values, err := parseFloat32Slice(data, fields)
	if err != nil {
		return MixerPerspective{}, err
	}

	return MixerPerspective{
		TopLeftX:     values[0],
		TopLeftY:     values[1],
		TopRightX:    values[2],
		TopRightY:    values[3],
		BottomLeftX:  values[4],
		BottomLeftY:  values[5],
		BottomRightX: values[6],
		BottomRightY: values[7],
	}, nil
}
