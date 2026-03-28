package commands_test

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

func TestBasicCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "LOG LEVEL",
			command: commands.DirectCommandLogLevel{
				Level: types.LogLevelInfo,
			},
			expected: "LOG LEVEL INFO",
		},
		{
			name: "PING without token",
			command: commands.DirectCommandPing{
				Token: nil,
			},
			expected: "PING",
		},
		{
			name: "PING with token",
			command: commands.DirectCommandPing{
				Token: ptr("abc123"),
			},
			expected: "PING abc123",
		},
		{
			name:     "BYE",
			command:  commands.CommandBye{},
			expected: "BYE",
		},
		{
			name:     "KILL",
			command:  commands.CommandKill{},
			expected: "KILL",
		},
		{
			name:     "RESTART",
			command:  commands.CommandRestart{},
			expected: "RESTART",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.command.String()
			if got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
