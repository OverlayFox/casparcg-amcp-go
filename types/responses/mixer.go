package responses

import (
	"fmt"
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
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

func MixerBlendModeFromResponse(data []string) (types.BlendMode, error) {
	if len(data) < 1 {
		return "", fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	return types.ParseBlendMode(data[0])
}

func BoolFromResponse(data []string) (bool, error) {
	if len(data) < 1 {
		return false, fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	return data[0] == "1", nil
}

func FloatFromResponse(data []string) (float32, error) {
	if len(data) < 1 {
		return 0, fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	opacity, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return 0, fmt.Errorf("invalid opacity value: %w", err)
	}
	return float32(opacity), nil
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
	X float32
	Y float32
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
