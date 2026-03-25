package responses

import (
	"fmt"
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

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

func BlendModeFromResponse(data []string) (types.BlendMode, error) {
	if len(data) < 1 {
		return "", fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	return types.ParseBlendMode(data[0])
}
