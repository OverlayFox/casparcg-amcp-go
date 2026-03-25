package types

import (
	"fmt"
	"strconv"
)

type Fade struct {
	Duration int // in frames
	Tween    TweenType
}

type MixerParamsFill struct {
	X      *float32
	Y      *float32
	XScale *float32
	YScale *float32
}

type MixerInfoLevels struct {
	MinInput float32
	MaxInput float32

	Gamma float32

	MinOutput float32
	MaxOutput float32
}

func MixerInfoLevelsFromResponse(data []string) (MixerInfoLevels, error) {
	if len(data) < 5 {
		return MixerInfoLevels{}, fmt.Errorf("unexpected response length: got %d, expected at least 5", len(data))
	}

	minInput, err := strconv.ParseFloat(data[0], 32)
	if err != nil {
		return MixerInfoLevels{}, fmt.Errorf("invalid MinInput value: %w", err)
	}
	maxInput, err := strconv.ParseFloat(data[1], 32)
	if err != nil {
		return MixerInfoLevels{}, fmt.Errorf("invalid MaxInput value: %w", err)
	}
	gamma, err := strconv.ParseFloat(data[2], 32)
	if err != nil {
		return MixerInfoLevels{}, fmt.Errorf("invalid Gamma value: %w", err)
	}
	minOutput, err := strconv.ParseFloat(data[3], 32)
	if err != nil {
		return MixerInfoLevels{}, fmt.Errorf("invalid MinOutput value: %w", err)
	}
	maxOutput, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return MixerInfoLevels{}, fmt.Errorf("invalid MaxOutput value: %w", err)
	}

	return MixerInfoLevels{
		MinInput:  float32(minInput),
		MaxInput:  float32(maxInput),
		Gamma:     float32(gamma),
		MinOutput: float32(minOutput),
		MaxOutput: float32(maxOutput),
	}, nil
}

type MixerCrop struct {
	LeftEdge   *float32
	TopEdge    *float32
	RightEdge  *float32
	BottomEdge *float32
}
