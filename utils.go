package casparcg

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/responses"
)

func matchesToCINF(matches []string) (responses.CINF, error) {
	if matches == nil || len(matches) != 7 {
		return responses.CINF{}, fmt.Errorf("unexpected format for CINF response: %s", matches)
	}

	cinfSize, err := strconv.Atoi(strings.TrimSpace(matches[3]))
	if err != nil {
		return responses.CINF{}, fmt.Errorf("invalid file size in CINF response: %s", matches[3])
	}

	cinfLastModified, err := time.Parse("20060102150405", strings.TrimSpace(matches[4]))
	if err != nil {
		return responses.CINF{}, fmt.Errorf("invalid last modified date in CINF response: %s", matches[4])
	}

	cinfFrameCount, err := strconv.Atoi(strings.TrimSpace(matches[5]))
	if err != nil {
		return responses.CINF{}, fmt.Errorf("invalid frame count in CINF response: %s", matches[5])
	}

	cinfFrameRate, err := types.StringToFrameRate(strings.TrimSpace(matches[6]))
	if err != nil {
		return responses.CINF{}, fmt.Errorf("invalid frame rate in CINF response: %s", matches[6])
	}

	return responses.CINF{
		Filename:     strings.TrimSpace(matches[1]),
		Type:         types.MediaTypes(strings.TrimSpace(matches[2])),
		FileSize:     int64(cinfSize),
		LastModified: cinfLastModified,
		FrameCount:   cinfFrameCount,
		FrameRate:    cinfFrameRate,
	}, nil
}

//nolint:unparam // minValue may vary in future use cases.
func inRangeFloat(valueName string, value, minValue, maxValue float32) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("%w: %s must be between %f and %f", ErrValueOutOfRange, valueName, minValue, maxValue)
	}
	return nil
}

func inRangeInt(valueName string, value, minValue, maxValue int) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("%w: %s must be between %d and %d", ErrValueOutOfRange, valueName, minValue, maxValue)
	}
	return nil
}
