package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

func appendParams(cmd string, params *[]string) string {
	if params != nil && len(*params) > 0 {
		quotedParams := make([]string, len(*params))
		for i, p := range *params {
			quotedParams[i] = quote(p)
		}
		return cmd + " " + strings.Join(quotedParams, " ")
	}
	return cmd
}

func appendInt(cmd string, value *int) string {
	if value != nil {
		return cmd + " " + strconv.Itoa(*value)
	}
	return cmd
}

func appendFloat(cmd string, value *float32) string {
	if value != nil {
		return cmd + " " + fmt.Sprintf("%.4g", *value)
	}
	return cmd
}

func appendBool(cmd string, value *bool) string {
	if value != nil {
		if *value {
			return cmd + " 1"
		}
		return cmd + " 0"
	}
	return cmd
}

func appendQuotedString(cmd string, value *string) string {
	if value != nil {
		return cmd + " " + quote(*value)
	}
	return cmd
}

func appendString(cmd string, value *string) string {
	if value != nil {
		return cmd + " " + *value
	}
	return cmd
}

func appendDurationTween(cmd string, duration *int, tween *types.TweenType) string {
	if duration != nil {
		cmd += " " + strconv.Itoa(*duration)
	}
	if tween != nil {
		cmd += " " + tween.String()
	}
	return cmd
}

func quote(s string) string {
	escaped := strings.ReplaceAll(s, `"`, `\"`)
	return "\"" + escaped + "\""
}

func baseCommand(command string, videoChannel int, layer *int) string {
	cmd := fmt.Sprintf("%s %d", command, videoChannel)
	if layer != nil {
		cmd += fmt.Sprintf("-%d", *layer)
	}
	return cmd
}

func ptr[T any](v T) *T {
	return &v
}
