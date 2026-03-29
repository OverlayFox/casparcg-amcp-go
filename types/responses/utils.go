package responses

import (
	"fmt"
	"strconv"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

// BoolFromResponse parses a boolean value from the response data.
func BoolFromResponse(data []string) (bool, error) {
	if len(data) < 1 {
		return false, fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	return data[0] == "1", nil
}

// IntFromResponse parses an integer value from the response data.
func IntFromResponse(data []string) (int, error) {
	if len(data) < 1 {
		return 0, fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	value, err := strconv.Atoi(data[0])
	if err != nil {
		return 0, fmt.Errorf("invalid integer value: %w", err)
	}
	return value, nil
}

// FloatFromResponse parses a float32 value from the response data.
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

// BlendModeFromResponse parses a blend mode from the response data.
func BlendModeFromResponse(data []string) (types.BlendMode, error) {
	if len(data) < 1 {
		return "", fmt.Errorf("unexpected response length: got %d, expected at least 1", len(data))
	}
	return types.ParseBlendMode(data[0])
}

// parseBool converts "1" to true, anything else to false.
func parseBool(value string) bool {
	return value == "1"
}

// parseFloat32 is a helper to parse a string to float32 with context.
func parseFloat32(value, fieldName string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid %s value: %w", fieldName, err)
	}
	return float32(f), nil
}

// parseFloat32Slice parses multiple float values from string slice.
func parseFloat32Slice(data []string, fieldNames []string) ([]float32, error) {
	if len(data) < len(fieldNames) {
		return nil, fmt.Errorf("unexpected response length: got %d, expected at least %d", len(data), len(fieldNames))
	}

	result := make([]float32, len(fieldNames))
	for i, name := range fieldNames {
		val, err := parseFloat32(data[i], name)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}
