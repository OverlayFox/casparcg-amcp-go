package commands_test

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

func TestLayerCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "LOAD with parameters",
			command: commands.LayerLoad{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Clip: "myclip",
				Parameters: &[]string{
					"param1",
					"param2",
				},
			},
			expected: "LOAD 1-2 \"myclip\" \"param1\" \"param2\"",
		},
		{
			name: "LOAD without parameters",
			command: commands.LayerLoad{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Clip: "myclip",
			},
			expected: "LOAD 1-2 \"myclip\"",
		},
		{
			name: "PLAY with clip and parameters",
			command: commands.LayerPlay{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Clip: ptr("myclip"),
				Parameters: &[]string{
					"param1",
					"param2",
				},
			},
			expected: "PLAY 1-2 myclip \"param1\" \"param2\"",
		},
		{
			name: "PLAY without clip and parameters",
			command: commands.LayerPlay{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "PLAY 1-2",
		},
		{
			name: "PAUSE",
			command: commands.LayerPause{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "PAUSE 1-2",
		},
		{
			name: "RESUME",
			command: commands.LayerResume{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "RESUME 1-2",
		},
		{
			name: "STOP",
			command: commands.LayerStop{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "STOP 1-2",
		},
		{
			name: "CLEAR on layer",
			command: commands.LayerClear{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "CLEAR 1-2",
		},
		{
			name: "CLEAR on channel",
			command: commands.LayerClear{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        nil,
				},
			},
			expected: "CLEAR 1",
		},
		{
			name: "CALL",
			command: commands.LayerCall{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Params: []string{"param1", "param2"},
			},
			expected: "CALL 1-2 \"param1\" \"param2\"",
		},
		{
			name: "SWAP layers with Transform",
			command: commands.LayerSwap{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				VideoChannel2: 3,
				Layer2:        ptr(4),
				Transform:     true,
			},
			expected: "SWAP 1-2 3-4 TRANSFORMS",
		},
		{
			name: "SWAP layers without Transform",
			command: commands.LayerSwap{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				VideoChannel2: 3,
				Layer2:        ptr(4),
				Transform:     false,
			},
			expected: "SWAP 1-2 3-4",
		},
		{
			name: "SWAP channels with Transform",
			command: commands.LayerSwap{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        nil,
				},
				VideoChannel2: 3,
				Layer2:        nil,
				Transform:     true,
			},
			expected: "SWAP 1 3 TRANSFORMS",
		},
		{
			name: "SWAP channels without Transform",
			command: commands.LayerSwap{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        nil,
				},
				VideoChannel2: 3,
				Layer2:        nil,
				Transform:     false,
			},
			expected: "SWAP 1 3",
		},
		{
			name: "ADD consumer with IDX and parameters",
			command: commands.LayerAdd{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx:  ptr(5),
				Params:       &[]string{"param1", "param2"},
				ConsumerName: "DECKLINK",
			},
			expected: "ADD 1-5 DECKLINK \"param1\" \"param2\"",
		},
		{
			name: "ADD consumer without IDX and parameters",
			command: commands.LayerAdd{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx:  nil,
				Params:       nil,
				ConsumerName: "DECKLINK",
			},
			expected: "ADD 1 DECKLINK",
		},
		{
			name: "ADD consumer without IDX but with parameters",
			command: commands.LayerAdd{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx:  nil,
				Params:       &[]string{"param1", "param2"},
				ConsumerName: "DECKLINK",
			},
			expected: "ADD 1 DECKLINK \"param1\" \"param2\"",
		},
		{
			name: "ADD consumer with IDX but without parameters",
			command: commands.LayerAdd{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx:  ptr(5),
				Params:       nil,
				ConsumerName: "DECKLINK",
			},
			expected: "ADD 1-5 DECKLINK",
		},
		{
			name: "REMOVE consumer with IDX",
			command: commands.LayerRemove{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx: ptr(5),
			},
			expected: "REMOVE 1-5",
		},
		{
			name: "REMOVE consumer without IDX but with parameters",
			command: commands.LayerRemove{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx: nil,
				Parameters:  &[]string{"param1", "param2"},
			},
			expected: "REMOVE 1 \"param1\" \"param2\"",
		},
		{
			name: "REMOVE consumer with IDX and with parameters", // IDX takes precedence over parameters
			command: commands.LayerRemove{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				ConsumerIdx: ptr(5),
				Parameters:  &[]string{"param1", "param2"},
			},
			expected: "REMOVE 1-5",
		},
		{
			name: "PRINT",
			command: commands.LayerPrint{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "PRINT 1",
		},
		{
			name: "SET mode with parameters",
			command: commands.LayerSet{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				VariableName: types.SetVariableMode,
				Value:        "VALUE",
			},
			expected: "SET 1 MODE VALUE",
		},
		{
			name: "LOCK ACQUIRE",
			command: commands.LayerLock{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Action:     types.LockActionAcquire,
				Passphrase: ptr("mypassword"),
			},
			expected: "LOCK 1 ACQUIRE mypassword",
		},
		{
			name: "LOCK RELEASE",
			command: commands.LayerLock{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Action:     types.LockActionRelease,
				Passphrase: ptr("mypassword"),
			},
			expected: "LOCK 1 RELEASE mypassword",
		},
		{
			name: "LOCK CLEAR",
			command: commands.LayerLock{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Action:     types.LockActionClear,
				Passphrase: nil,
			},
			expected: "LOCK 1 CLEAR",
		},
		{
			name: "INFO on layer",
			command: commands.LayerInfo{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "INFO 1-2",
		},
		{
			name: "INFO on channel",
			command: commands.LayerInfo{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        nil,
				},
			},
			expected: "INFO 1",
		},
		{
			name: "INFO on layer",
			command: commands.LayerInfoDelay{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "INFO 1-2 DELAY",
		},
		{
			name: "INFO on channel",
			command: commands.LayerInfoDelay{
				LayerCommand: commands.LayerCommand{
					VideoChannel: 1,
					Layer:        nil,
				},
			},
			expected: "INFO 1 DELAY",
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
