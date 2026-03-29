package commands_test

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

func TestQueryCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "CINF",
			command: commands.QueryCINF{
				Filename: "video.mp4",
			},
			expected: `CINF "video.mp4"`,
		},
		{
			name: "CLS with directory",
			command: commands.QueryCLS{
				Directory: ptr("subdir"),
			},
			expected: `CLS "subdir"`,
		},
		{
			name:     "CLS without directory",
			command:  commands.QueryCLS{},
			expected: "CLS",
		},
		{
			name:     "FLS",
			command:  commands.QueryFLS{},
			expected: "FLS",
		},
		{
			name: "TLS with directory",
			command: commands.QueryTLS{
				Directory: ptr("templates"),
			},
			expected: "TLS templates",
		},
		{
			name:     "TLS without directory",
			command:  commands.QueryTLS{},
			expected: "TLS",
		},
		{
			name: "VERSION without component",
			command: commands.QueryVersion{
				Component: "",
			},
			expected: "VERSION",
		},
		{
			name: "VERSION with component",
			command: commands.QueryVersion{
				Component: types.VersionInfoServer,
			},
			expected: "VERSION SERVER",
		},
		{
			name: "INFO without component",
			command: commands.QueryInfo{
				Component: "",
			},
			expected: "INFO",
		},
		{
			name: "INFO with component",
			command: commands.QueryInfo{
				Component: types.InfoComponentConfig,
			},
			expected: "INFO CONFIG",
		},
		{
			name: "INFO TEMPLATE",
			command: commands.QueryInfoTemplate{
				Template: "L3",
			},
			expected: "INFO TEMPLATE \"L3\"",
		},
		{
			name:     "DIAG",
			command:  commands.QueryDiag{},
			expected: "DIAG",
		},
		{
			name:     "GL INFO",
			command:  commands.QueryGLInfo{},
			expected: "GL INFO",
		},
		{
			name:     "GL GC",
			command:  commands.QueryGLGC{},
			expected: "GL GC",
		},
		{
			name:     "HELP",
			command:  commands.QueryHelp{},
			expected: "HELP",
		},
		{
			name: "HELP PRODUCER with argument",
			command: commands.QueryHelpProducer{
				Producer: ptr("AMCP"),
			},
			expected: "HELP PRODUCER AMCP",
		},
		{
			name: "HELP PRODUCER without argument",
			command: commands.QueryHelpProducer{
				Producer: nil,
			},
			expected: "HELP PRODUCER",
		},
		{
			name: "HELP CONSUMER with argument",
			command: commands.QueryHelpConsumer{
				Consumer: ptr("AMCP"),
			},
			expected: "HELP CONSUMER AMCP",
		},
		{
			name: "HELP CONSUMER without argument",
			command: commands.QueryHelpConsumer{
				Consumer: nil,
			},
			expected: "HELP CONSUMER",
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
