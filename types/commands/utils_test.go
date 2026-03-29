//nolint:testpackage // Testing internal/unexported functions
package commands

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

func TestUtilsEmpty(t *testing.T) {
	testNilling := []struct {
		name     string
		expected string
		function func() string
	}{
		{
			name:     "appendParams - with empty",
			expected: "TEST",
			function: func() string {
				return appendParams("TEST", &[]string{})
			},
		},
		{
			name:     "appendInt - with empty",
			expected: "TEST",
			function: func() string {
				return appendInt("TEST", nil)
			},
		},
		{
			name:     "appendFloat - with empty",
			expected: "TEST",
			function: func() string {
				return appendFloat("TEST", nil)
			},
		},
		{
			name:     "appendBool - with empty",
			expected: "TEST",
			function: func() string {
				return appendBool("TEST", nil)
			},
		},
		{
			name:     "appendBool - with empty",
			expected: "TEST",
			function: func() string {
				return appendBool("TEST", nil)
			},
		},
		{
			name:     "appendQuotedString - with empty",
			expected: "TEST",
			function: func() string {
				return appendQuotedString("TEST", nil)
			},
		},
		{
			name:     "appendString - with empty",
			expected: "TEST",
			function: func() string {
				return appendString("TEST", nil)
			},
		},
		{
			name:     "appendDurationTween - with empty",
			expected: "TEST",
			function: func() string {
				return appendDurationTween("TEST", nil, nil)
			},
		},
	}

	for _, tt := range testNilling {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestUtilsNonEmpty(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		function func() string
	}{
		{
			name:     "appendParams - with data",
			expected: "TEST \"param1\" \"param2\"",
			function: func() string {
				params := []string{"param1", "param2"}
				return appendParams("TEST", &params)
			},
		},
		{
			name:     "appendParams - with json",
			expected: `TEST "{\"key\": \"value\"}" "{\"key2\": \"value2\"}"`,
			function: func() string {
				params := []string{`{"key": "value"}`, `{"key2": "value2"}`}
				return appendParams("TEST", &params)
			},
		},
		{
			name:     "appendParams - with xml data",
			expected: "TEST \"<data>value</data>\" \"<data2>value2</data2>\"",
			function: func() string {
				params := []string{`<data>value</data>`, `<data2>value2</data2>`}
				return appendParams("TEST", &params)
			},
		},
		{
			name:     "appendInt - with data",
			expected: "TEST 42",
			function: func() string {
				value := 42
				return appendInt("TEST", &value)
			},
		},
		{
			name:     "appendFloat - with data",
			expected: "TEST 3.142",
			function: func() string {
				value := float32(3.14159)
				return appendFloat("TEST", &value)
			},
		},
		{
			name:     "appendBool - with true",
			expected: "TEST 1",
			function: func() string {
				value := true
				return appendBool("TEST", &value)
			},
		},
		{
			name:     "appendBool - with false",
			expected: "TEST 0",
			function: func() string {
				value := false
				return appendBool("TEST", &value)
			},
		},
		{
			name:     "appendQuotedString - with data",
			expected: "TEST \"quoted string\"",
			function: func() string {
				value := "quoted string"
				return appendQuotedString("TEST", &value)
			},
		},
		{
			name:     "appendString - with data",
			expected: "TEST simple string",
			function: func() string {
				value := "simple string"
				return appendString("TEST", &value)
			},
		},
		{
			name:     "appendDurationTween - with data",
			expected: "TEST 500 easeinsine",
			function: func() string {
				duration := 500
				tween := types.TweenTypeEaseInSine
				return appendDurationTween("TEST", &duration, &tween)
			},
		},
		{
			name:     "baseCommand - with layer",
			expected: "TEST 1-2",
			function: func() string {
				return baseCommand("TEST", 1, ptr(2))
			},
		},
		{
			name:     "baseCommand - without layer",
			expected: "TEST 1",
			function: func() string {
				return baseCommand("TEST", 1, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
