package casparcg

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

func TestCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "CG STOP",
			command: types.TemplateCommandCGStop{
				TemplateCommandCG: types.TemplateCommandCG{
					VideoChannel: 1,
					Layer:        12,
				},
				CgLayer: 2,
			},
			expected: "CG 1-12 STOP 2",
		},
		{
			name: "PLAY with clip",
			command: types.CommandPlay{
				BasicCommand: types.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
				Clip: strPtr("AMB"),
			},
			expected: `PLAY 1-10 "AMB"`,
		},
		{
			name: "PLAY without clip",
			command: types.CommandPlay{
				BasicCommand: types.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "PLAY 1-10",
		},
		{
			name: "LOAD",
			command: types.CommandLoad{
				BasicCommand: types.BasicCommand{
					VideoChannel: 1,
					Layer:        11,
				},
				Clip: "myclip",
			},
			expected: `LOAD 1-11 "myclip"`,
		},
		{
			name: "PAUSE",
			command: types.CommandPause{
				BasicCommand: types.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "PAUSE 1-10",
		},
		{
			name: "STOP",
			command: types.CommandStop{
				BasicCommand: types.BasicCommand{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "STOP 1-10",
		},
		{
			name: "CLS without directory",
			command: types.QueryCommandCLS{
				Directory: nil,
			},
			expected: "CLS",
		},
		{
			name: "CLS with directory",
			command: types.QueryCommandCLS{
				Directory: strPtr("subfolder"),
			},
			expected: `CLS "subfolder"`,
		},
		{
			name:     "VERSION",
			command:  types.QueryCommandVersion{},
			expected: "VERSION",
		},
		{
			name: "CG ADD with data",
			command: types.TemplateCommandCGAdd{
				TemplateCommandCG: types.TemplateCommandCG{
					VideoChannel: 1,
					Layer:        10,
				},
				CgLayer:    1,
				Template:   "lower_third",
				PlayOnLoad: true,
				Data:       strPtr(`{"f0":"Hello"}`),
			},
			expected: "CG 1-10 ADD 1 \"lower_third\" 1 \"{\"f0\":\"Hello\"}\"",
		},
		{
			name: "CG CLEAR",
			command: types.TemplateCommandCGClear{
				TemplateCommandCG: types.TemplateCommandCG{
					VideoChannel: 1,
					Layer:        10,
				},
			},
			expected: "CG 1-10 CLEAR",
		},
		{
			name: "LOGLEVEL",
			command: types.CommandLogLevel{
				Level: types.LogLevelInfo,
			},
			expected: "LOG LEVEL info",
		},
		{
			name: "SWAP layers",
			command: types.CommandSwap{
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
			command: types.CommandSwap{
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
