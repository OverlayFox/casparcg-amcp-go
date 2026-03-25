package responses

import (
	"fmt"
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type MixerInfoChroma struct {
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

func MixerInfoChromaFromResponse(data []string) (MixerInfoChroma, error) {
	if len(data) < 9 {
		return MixerInfoChroma{}, fmt.Errorf("unexpected response length: got %d, expected at least 9", len(data))
	}

	enabled := data[0] == "1"
	targetHue, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid TargetHue value: %w", err)
	}
	hueWidth, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid HueWidth value: %w", err)
	}
	minSaturation, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid MinSaturation value: %w", err)
	}
	minBrightness, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid MinBrightness value: %w", err)
	}
	softness, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid Softness value: %w", err)
	}
	spillSuppress, err := strconv.ParseFloat(data[6], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid SpillSuppress value: %w", err)
	}
	spillSuppressSaturation, err := strconv.ParseFloat(data[7], 32)
	if err != nil {
		return MixerInfoChroma{}, fmt.Errorf("invalid SpillSuppressSaturation value: %w", err)
	}
	showMask := data[8] == "1"

	return MixerInfoChroma{
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

type MixerInfoFill struct {
	X      float32
	Y      float32
	XScale float32
	YScale float32
}

func MixerInfoFillFromResponse(data []string) (MixerInfoFill, error) {
	if len(data) < 4 {
		return MixerInfoFill{}, fmt.Errorf("unexpected response length: got %d, expected at least 4", len(data))
	}

	x, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerInfoFill{}, fmt.Errorf("invalid X value: %w", err)
	}
	y, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerInfoFill{}, fmt.Errorf("invalid Y value: %w", err)
	}
	xScale, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerInfoFill{}, fmt.Errorf("invalid XScale value: %w", err)
	}
	yScale, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerInfoFill{}, fmt.Errorf("invalid YScale value: %w", err)
	}

	return MixerInfoFill{
		X:      float32(x),
		Y:      float32(y),
		XScale: float32(xScale),
		YScale: float32(yScale),
	}, nil
}
