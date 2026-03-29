package commands_test

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

func TestTemplateCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "CG ADD with play on load and data",
			command: commands.TemplateCGAdd{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(1),
				},
				Template:   "template1",
				PlayOnLoad: true,
				Data:       ptr(`{"key":"value"}`),
			},
			expected: `CG 1-9999 ADD 1 "template1" 1 "{\"key\":\"value\"}"`,
		},
		{
			name: "CG ADD without play on load and data",
			command: commands.TemplateCGAdd{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(1),
				},
				Template:   "template1",
				PlayOnLoad: false,
				Data:       ptr(`{"key":"value"}`),
			},
			expected: `CG 1-9999 ADD 1 "template1" 0 "{\"key\":\"value\"}"`,
		},
		{
			name: "CG ADD without play on load and without data",
			command: commands.TemplateCGAdd{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(1),
				},
				Template:   "template1",
				PlayOnLoad: false,
				Data:       nil,
			},
			expected: `CG 1-9999 ADD 1 "template1" 0`,
		},
		{
			name: "CG PLAY",
			command: commands.TemplateCGPlay{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
			},
			expected: `CG 1-9999 PLAY 2`,
		},
		{
			name: "CG STOP",
			command: commands.TemplateCGStop{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
			},
			expected: `CG 1-9999 STOP 2`,
		},
		{
			name: "CG NEXT",
			command: commands.TemplateCGNext{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
			},
			expected: `CG 1-9999 NEXT 2`,
		},
		{
			name: "CG REMOVE",
			command: commands.TemplateCGRemove{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
			},
			expected: `CG 1-9999 REMOVE 2`,
		},
		{
			name: "CG CLEAR",
			command: commands.TemplateCGClear{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
			},
			expected: `CG 1-9999 CLEAR`,
		},
		{
			name: "CG UPDATE",
			command: commands.TemplateCGUpdate{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
				Data: `{"key":"value"}`,
			},
			expected: `CG 1-9999 UPDATE 2 "{\"key\":\"value\"}"`,
		},
		{
			name: "CG INVOKE",
			command: commands.TemplateCGInvoke{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
				Method: "methodName",
			},
			expected: `CG 1-9999 INVOKE 2 "methodName"`,
		},
		{
			name: "CG INFO",
			command: commands.TemplateCGInfo{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      ptr(2),
				},
			},
			expected: `CG 1-9999 INFO 2`,
		},
		{
			name: "CG INFO without CG layer",
			command: commands.TemplateCGInfo{
				CGCommand: commands.CGCommand{
					VideoChannel: 1,
					Layer:        ptr(9999),
					CgLayer:      nil,
				},
			},
			expected: `CG 1-9999 INFO`,
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
