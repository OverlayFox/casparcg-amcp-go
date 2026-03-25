package responses

import (
	"fmt"
	"strconv"
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
	if len(data) < 9 {
		return MixerChroma{}, fmt.Errorf("unexpected response length: got %d, expected at least 9", len(data))
	}

	enabled := data[0] == "1"
	targetHue, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid TargetHue value: %w", err)
	}
	hueWidth, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid HueWidth value: %w", err)
	}
	minSaturation, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid MinSaturation value: %w", err)
	}
	minBrightness, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid MinBrightness value: %w", err)
	}
	softness, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid Softness value: %w", err)
	}
	spillSuppress, err := strconv.ParseFloat(data[6], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid SpillSuppress value: %w", err)
	}
	spillSuppressSaturation, err := strconv.ParseFloat(data[7], 32)
	if err != nil {
		return MixerChroma{}, fmt.Errorf("invalid SpillSuppressSaturation value: %w", err)
	}
	showMask := data[8] == "1"

	return MixerChroma{
		Enabled:                 enabled,
		TargetHue:               float32(targetHue),
		HueWidth:                float32(hueWidth),
		MinSaturation:           float32(minSaturation),
		MinBrightness:           float32(minBrightness),
		Softness:                float32(softness),
		SpillSuppress:           float32(spillSuppress),
		SpillSuppressSaturation: float32(spillSuppressSaturation),
		ShowMask:                showMask,
	}, nil
}

type MixerFill struct {
	X      float32
	Y      float32
	XScale float32
	YScale float32
}

func MixerFillFromResponse(data []string) (MixerFill, error) {
	if len(data) < 4 {
		return MixerFill{}, fmt.Errorf("unexpected response length: got %d, expected at least 4", len(data))
	}

	x, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerFill{}, fmt.Errorf("invalid X value: %w", err)
	}
	y, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerFill{}, fmt.Errorf("invalid Y value: %w", err)
	}
	xScale, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerFill{}, fmt.Errorf("invalid XScale value: %w", err)
	}
	yScale, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerFill{}, fmt.Errorf("invalid YScale value: %w", err)
	}

	return MixerFill{
		X:      float32(x),
		Y:      float32(y),
		XScale: float32(xScale),
		YScale: float32(yScale),
	}, nil
}

type MixerClip struct {
	X      float32 // X: The new x position, 0 = left edge of monitor, 0.5 = middle of monitor, 1.0 = right edge of monitor. Higher and lower values allowed.
	Y      float32 // Y: The new y position, 0 = top edge of monitor, 0.5 = middle of monitor, 1.0 = bottom edge of monitor. Higher and lower values allowed.
	Width  float32 // Width: The new width, 1 = 1x the screen width, 0.5 = half the screen width. Higher and lower values allowed. Negative values flips the layer.
	Height float32 // Height: The new height, 1 = 1x the screen height, 0.5 = half the screen height. Higher and lower values allowed. Negative values flips the layer.
}

func MixerClipFromResponse(data []string) (MixerClip, error) {
	if len(data) < 4 {
		return MixerClip{}, fmt.Errorf("unexpected response length: got %d, expected at least 4", len(data))
	}

	x, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerClip{}, fmt.Errorf("invalid X value: %w", err)
	}
	y, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerClip{}, fmt.Errorf("invalid Y value: %w", err)
	}
	width, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerClip{}, fmt.Errorf("invalid Width value: %w", err)
	}
	height, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerClip{}, fmt.Errorf("invalid Height value: %w", err)
	}

	return MixerClip{
		X:      float32(x),
		Y:      float32(y),
		Width:  float32(width),
		Height: float32(height),
	}, nil
}

type MixerAnchor struct {
	X float32 // X defines the x anchor point, 0 = left edge of layer, 0.5 = middle of layer, 1.0 = right edge of layer. Higher and lower values allowed.
	Y float32 // Y defines the y anchor point, 0 = top edge of layer, 0.5 = middle of layer, 1.0 = bottom edge of layer. Higher and lower values allowed.
}

func MixerAnchorFromResponse(data []string) (MixerAnchor, error) {
	if len(data) < 2 {
		return MixerAnchor{}, fmt.Errorf("unexpected response length: got %d, expected at least 2", len(data))
	}

	x, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerAnchor{}, fmt.Errorf("invalid X value: %w", err)
	}
	y, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerAnchor{}, fmt.Errorf("invalid Y value: %w", err)
	}

	return MixerAnchor{
		X: float32(x),
		Y: float32(y),
	}, nil
}

type MixerCrop struct {
	LeftEdge   float32
	TopEdge    float32
	RightEdge  float32
	BottomEdge float32
}

func MixerCropFromResponse(data []string) (MixerCrop, error) {
	if len(data) < 4 {
		return MixerCrop{}, fmt.Errorf("unexpected response length: got %d, expected at least 4", len(data))
	}

	leftEdge, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerCrop{}, fmt.Errorf("invalid LeftEdge value: %w", err)
	}
	topEdge, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerCrop{}, fmt.Errorf("invalid TopEdge value: %w", err)
	}
	rightEdge, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerCrop{}, fmt.Errorf("invalid RightEdge value: %w", err)
	}
	bottomEdge, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerCrop{}, fmt.Errorf("invalid BottomEdge value: %w", err)
	}

	return MixerCrop{
		LeftEdge:   float32(leftEdge),
		TopEdge:    float32(topEdge),
		RightEdge:  float32(rightEdge),
		BottomEdge: float32(bottomEdge),
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
	if len(data) < 8 {
		return MixerPerspective{}, fmt.Errorf("unexpected response length: got %d, expected at least 8", len(data))
	}

	topLeftX, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid TopLeftX value: %w", err)
	}
	topLeftY, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid TopLeftY value: %w", err)
	}
	topRightX, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid TopRightX value: %w", err)
	}
	topRightY, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid TopRightY value: %w", err)
	}
	bottomLeftX, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid BottomLeftX value: %w", err)
	}
	bottomLeftY, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid BottomLeftY value: %w", err)
	}
	bottomRightX, err := strconv.ParseFloat(data[6], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid BottomRightX value: %w", err)
	}
	bottomRightY, err := strconv.ParseFloat(data[7], 32)
	if err != nil {
		return MixerPerspective{}, fmt.Errorf("invalid BottomRightY value: %w", err)
	}

	return MixerPerspective{
		TopLeftX:     float32(topLeftX),
		TopLeftY:     float32(topLeftY),
		TopRightX:    float32(topRightX),
		TopRightY:    float32(topRightY),
		BottomLeftX:  float32(bottomLeftX),
		BottomLeftY:  float32(bottomLeftY),
		BottomRightX: float32(bottomRightX),
		BottomRightY: float32(bottomRightY),
	}, nil
}
