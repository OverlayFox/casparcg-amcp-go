//go:build !integration

package casparcg_test

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

func TestCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "CG STOP",
			command: commands.TemplateCommandCGStop{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        intPtr(12),
					CgLayer:      intPtr(2),
				},
			},
			expected: "CG 1-12 STOP 2",
		},
		{
			name: "PLAY with clip",
			command: commands.CommandPlay{
				BasicCommand: commands.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
				Clip: strPtr("AMB"),
			},
			expected: `PLAY 1-10 "AMB"`,
		},
		{
			name: "PLAY without clip",
			command: commands.CommandPlay{
				BasicCommand: commands.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "PLAY 1-10",
		},
		{
			name: "LOAD",
			command: commands.CommandLoad{
				BasicCommand: commands.BasicCommand{
					VideoChannel: 1,
					Layer:        11,
				},
				Clip: "myclip",
			},
			expected: `LOAD 1-11 "myclip"`,
		},
		{
			name: "PAUSE",
			command: commands.CommandPause{
				BasicCommand: commands.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "PAUSE 1-10",
		},
		{
			name: "STOP",
			command: commands.CommandStop{
				BasicCommand: commands.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "STOP 1-10",
		},
		{
			name: "CLS without directory",
			command: commands.QueryCommandCLS{
				Directory: nil,
			},
			expected: "CLS",
		},
		{
			name: "CLS with directory",
			command: commands.QueryCommandCLS{
				Directory: strPtr("subfolder"),
			},
			expected: `CLS "subfolder"`,
		},
		{
			name:     "VERSION",
			command:  commands.QueryCommandVersion{},
			expected: "VERSION",
		},
		{
			name: "CG ADD with data",
			command: commands.TemplateCommandCGAdd{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        intPtr(10),
					CgLayer:      intPtr(1),
				},
				Template:   "lower_third",
				PlayOnLoad: true,
				Data:       strPtr(`{"f0":"Hello"}`),
			},
			expected: "CG 1-10 ADD 1 \"lower_third\" 1 \"{\\\"f0\\\":\\\"Hello\\\"}\"",
		},
		{
			name: "CG CLEAR",
			command: commands.TemplateCommandCGClear{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        intPtr(10),
				},
			},
			expected: "CG 1-10 CLEAR",
		},
		{
			name: "LOGLEVEL",
			command: commands.CommandLogLevel{
				Level: types.LogLevelInfo,
			},
			expected: "LOG LEVEL info",
		},
		{
			name: "SWAP layers",
			command: commands.CommandSwap{
				VideoChannel1: 1,
				Layer1:        intPtr(10),
				VideoChannel2: 1,
				Layer2:        intPtr(20),
				Transform:     false,
			},
			expected: "SWAP 1-10 1-20",
		},
		{
			name: "SWAP with transforms",
			command: commands.CommandSwap{
				VideoChannel1: 1,
				Layer1:        intPtr(10),
				VideoChannel2: 2,
				Layer2:        intPtr(20),
				Transform:     true,
			},
			expected: "SWAP 1-10 2-20 TRANSFORMS",
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

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
