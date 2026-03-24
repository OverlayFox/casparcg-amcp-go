package returns

import (
	"fmt"
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type MixerChromaInfo struct {
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

func MixerChromaInfoFromResponse(data []string) (MixerChromaInfo, error) {
	if len(data) < 9 {
		return MixerChromaInfo{}, fmt.Errorf("unexpected response length: got %d, expected at least 9", len(data))
	}

	enabled := data[0] == "1"
	targetHue, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid TargetHue value: %w", err)
	}
	hueWidth, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid HueWidth value: %w", err)
	}
	minSaturation, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid MinSaturation value: %w", err)
	}
	minBrightness, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid MinBrightness value: %w", err)
	}
	softness, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid Softness value: %w", err)
	}
	spillSuppress, err := strconv.ParseFloat(data[6], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid SpillSuppress value: %w", err)
	}
	spillSuppressSaturation, err := strconv.ParseFloat(data[7], 32)
	if err != nil {
		return MixerChromaInfo{}, fmt.Errorf("invalid SpillSuppressSaturation value: %w", err)
	}
	showMask := data[8] == "1"

	return MixerChromaInfo{
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

type MixerLevelsInfo struct {
	MinInput float32
	MaxInput float32

	Gamma float32

	MinOutput float32
	MaxOutput float32
}

func MixerLevelsInfoFromResponse(data []string) (MixerLevelsInfo, error) {
	if len(data) < 5 {
		return MixerLevelsInfo{}, fmt.Errorf("unexpected response length: got %d, expected at least 5", len(data))
	}

	minInput, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerLevelsInfo{}, fmt.Errorf("invalid MinInput value: %w", err)
	}
	maxInput, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerLevelsInfo{}, fmt.Errorf("invalid MaxInput value: %w", err)
	}
	gamma, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerLevelsInfo{}, fmt.Errorf("invalid Gamma value: %w", err)
	}
	minOutput, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerLevelsInfo{}, fmt.Errorf("invalid MinOutput value: %w", err)
	}
	maxOutput, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerLevelsInfo{}, fmt.Errorf("invalid MaxOutput value: %w", err)
	}

	return MixerLevelsInfo{
		MinInput:  float32(minInput),
		MaxInput:  float32(maxInput),
		Gamma:     float32(gamma),
		MinOutput: float32(minOutput),
		MaxOutput: float32(maxOutput),
	}, nil
}
