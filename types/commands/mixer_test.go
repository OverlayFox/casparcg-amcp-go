package commands_test

import (
	"testing"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

func TestMixerCommandSerialization(t *testing.T) {
	tests := []struct {
		name     string
		command  interface{ String() string }
		expected string
	}{
		{
			name: "Get MASTERVOLUME",
			command: commands.MixerMasterVolume{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
			},
			expected: "MIXER 1 MASTERVOLUME",
		},
		{
			name: "Set MASTERVOLUME",
			command: commands.MixerMasterVolume{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
				Volume: ptr(float32(0.5)),
			},
			expected: "MIXER 1 MASTERVOLUME 0.5",
		},
		{
			name: "Get STRAIGHT ALPHA OUTPUT",
			command: commands.MixerStraightAlphaOutput{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
			},
			expected: "MIXER 1 STRAIGHT_ALPHA_OUTPUT",
		},
		{
			name: "Set STRAIGHT ALPHA OUTPUT true",
			command: commands.MixerStraightAlphaOutput{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
				Enable: ptr(true),
			},
			expected: "MIXER 1 STRAIGHT_ALPHA_OUTPUT 1",
		},
		{
			name: "Set STRAIGHT ALPHA OUTPUT false",
			command: commands.MixerStraightAlphaOutput{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
				Enable: ptr(false),
			},
			expected: "MIXER 1 STRAIGHT_ALPHA_OUTPUT 0",
		},
		{
			name: "Get GRID",
			command: commands.MixerGrid{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
			},
			expected: "MIXER 1 GRID",
		},
		{
			name: "Set GRID",
			command: commands.MixerGrid{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
				Resolution: ptr(2),
			},
			expected: "MIXER 1 GRID 2",
		},
		{
			name: "Set GRID with fade",
			command: commands.MixerGrid{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
				Resolution: ptr(2),
				Duration:   ptr(50),
				Tween:      ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1 GRID 2 50 easeinsine",
		},
		{
			name: "COMMIT",
			command: commands.MixerCommit{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
			},
			expected: "MIXER 1 COMMIT",
		},
		{
			name: "CLEAR Channel",
			command: commands.MixerClear{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
				},
			},
			expected: "MIXER 1 CLEAR",
		},
		{
			name: "CLEAR Layer",
			command: commands.MixerClear{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 CLEAR",
		},
		{
			name: "KEYER enabled",
			command: commands.MixerKeyer{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Show: true,
			},
			expected: "MIXER 1-2 KEYER 1",
		},
		{
			name: "KEYER disabled",
			command: commands.MixerKeyer{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Show: false,
			},
			expected: "MIXER 1-2 KEYER 0",
		},
		{
			name: "CHROMA info",
			command: commands.MixerChroma{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 CHROMA",
		},
		{
			name: "CHROMA set",
			command: commands.MixerChroma{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Enable:                  ptr(true),
				TargetHue:               ptr(float32(120)),
				HueWidth:                ptr(float32(0.1)),
				MinSaturation:           ptr(float32(0.2)),
				MinBrightness:           ptr(float32(0.3)),
				Softness:                ptr(float32(0.4)),
				SpillSuppress:           ptr(float32(0.5)),
				SpillSuppressSaturation: ptr(float32(0.6)),
				ShowMask:                ptr(true),
			},
			expected: "MIXER 1-2 CHROMA 1 120 0.1 0.2 0.3 0.4 0.5 0.6 1",
		},
		{
			name: "CHROMA set with tween",
			command: commands.MixerChroma{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Enable:                  ptr(true),
				TargetHue:               ptr(float32(120)),
				HueWidth:                ptr(float32(0.1)),
				MinSaturation:           ptr(float32(0.2)),
				MinBrightness:           ptr(float32(0.3)),
				Softness:                ptr(float32(0.4)),
				SpillSuppress:           ptr(float32(0.5)),
				SpillSuppressSaturation: ptr(float32(0.6)),
				ShowMask:                ptr(true),
				FadeDuration:            ptr(50),
				Tween:                   ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 CHROMA 1 120 0.1 0.2 0.3 0.4 0.5 0.6 1 50 easeinsine",
		},
		{
			name: "CHROMA unset",
			command: commands.MixerChroma{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Enable: ptr(false),
			},
			expected: "MIXER 1-2 CHROMA 0",
		},
		{
			name: "Get BLEND mode",
			command: commands.MixerBlend{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 BLEND",
		},
		{
			name: "Set BLEND mode",
			command: commands.MixerBlend{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				BlendMode: ptr(types.BlendModeScreen),
			},
			expected: "MIXER 1-2 BLEND SCREEN",
		},
		{
			name: "Get INVERT",
			command: commands.MixerInvert{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 INVERT",
		},
		{
			name: "Set INVERT enabled",
			command: commands.MixerInvert{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Invert: ptr(true),
			},
			expected: "MIXER 1-2 INVERT 1",
		},
		{
			name: "Set INVERT disabled",
			command: commands.MixerInvert{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Invert: ptr(false),
			},
			expected: "MIXER 1-2 INVERT 0",
		},
		{
			name: "Get OPACITY",
			command: commands.MixerOpacity{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 OPACITY",
		},
		{
			name: "Set OPACITY",
			command: commands.MixerOpacity{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Opacity: ptr(float32(0.5)),
			},
			expected: "MIXER 1-2 OPACITY 0.5",
		},
		{
			name: "Set OPACITY with fade",
			command: commands.MixerOpacity{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Opacity:  ptr(float32(0.5)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 OPACITY 0.5 50 easeinsine",
		},
		{
			name: "Get BRIGHTNESS",
			command: commands.MixerBrightness{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 BRIGHTNESS",
		},
		{
			name: "Set BRIGHTNESS",
			command: commands.MixerBrightness{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Brightness: ptr(float32(0.5)),
			},
			expected: "MIXER 1-2 BRIGHTNESS 0.5",
		},
		{
			name: "Set BRIGHTNESS with fade",
			command: commands.MixerBrightness{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Brightness: ptr(float32(0.5)),
				Duration:   ptr(50),
				Tween:      ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 BRIGHTNESS 0.5 50 easeinsine",
		},
		{
			name: "Get SATURATION",
			command: commands.MixerSaturation{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 SATURATION",
		},
		{
			name: "Set SATURATION",
			command: commands.MixerSaturation{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Saturation: ptr(float32(0.5)),
			},
			expected: "MIXER 1-2 SATURATION 0.5",
		},
		{
			name: "Set SATURATION with fade",
			command: commands.MixerSaturation{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Saturation: ptr(float32(0.5)),
				Duration:   ptr(50),
				Tween:      ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 SATURATION 0.5 50 easeinsine",
		},
		{
			name: "Get CONTRAST",
			command: commands.MixerContrast{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 CONTRAST",
		},
		{
			name: "Set CONTRAST",
			command: commands.MixerContrast{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Contrast: ptr(float32(0.5)),
			},
			expected: "MIXER 1-2 CONTRAST 0.5",
		},
		{
			name: "Set CONTRAST with fade",
			command: commands.MixerContrast{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Contrast: ptr(float32(0.5)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 CONTRAST 0.5 50 easeinsine",
		},
		{
			name: "Get LEVELS",
			command: commands.MixerLevels{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 LEVELS",
		},
		{
			name: "Set LEVELS",
			command: commands.MixerLevels{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				MinInput:  ptr(float32(0.0)),
				MaxInput:  ptr(float32(1.0)),
				Gamma:     ptr(float32(1.0)),
				MinOutput: ptr(float32(0.0)),
				MaxOutput: ptr(float32(1.0)),
			},
			expected: "MIXER 1-2 LEVELS 0 1 1 0 1",
		},
		{
			name: "Set LEVELS with fade",
			command: commands.MixerLevels{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				MinInput:  ptr(float32(0.0)),
				MaxInput:  ptr(float32(1.0)),
				Gamma:     ptr(float32(1.0)),
				MinOutput: ptr(float32(0.0)),
				MaxOutput: ptr(float32(1.0)),
				Duration:  ptr(50),
				Tween:     ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 LEVELS 0 1 1 0 1 50 easeinsine",
		},
		{
			name: "Get FILL",
			command: commands.MixerFill{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 FILL",
		},
		{
			name: "Set FILL",
			command: commands.MixerFill{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				X:      ptr(float32(0.0)),
				Y:      ptr(float32(0.1)),
				XScale: ptr(float32(0.2)),
				YScale: ptr(float32(0.3)),
			},
			expected: "MIXER 1-2 FILL 0 0.1 0.2 0.3",
		},
		{
			name: "Set FILL with fade",
			command: commands.MixerFill{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				X:        ptr(float32(0.0)),
				Y:        ptr(float32(0.1)),
				XScale:   ptr(float32(0.2)),
				YScale:   ptr(float32(0.3)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 FILL 0 0.1 0.2 0.3 50 easeinsine",
		},
		{
			name: "Get CLIP",
			command: commands.MixerClip{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 CLIP",
		},
		{
			name: "Set CLIP",
			command: commands.MixerClip{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				X:      ptr(float32(0.0)),
				Y:      ptr(float32(0.1)),
				Width:  ptr(float32(0.2)),
				Height: ptr(float32(0.3)),
			},
			expected: "MIXER 1-2 CLIP 0 0.1 0.2 0.3",
		},
		{
			name: "Set CLIP with fade",
			command: commands.MixerClip{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				X:        ptr(float32(0.0)),
				Y:        ptr(float32(0.1)),
				Width:    ptr(float32(0.2)),
				Height:   ptr(float32(0.3)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 CLIP 0 0.1 0.2 0.3 50 easeinsine",
		},
		{
			name: "Get ANCHOR",
			command: commands.MixerAnchor{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 ANCHOR",
		},
		{
			name: "Set ANCHOR",
			command: commands.MixerAnchor{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				X: ptr(float32(0.0)),
				Y: ptr(float32(0.1)),
			},
			expected: "MIXER 1-2 ANCHOR 0 0.1",
		},
		{
			name: "Set ANCHOR with fade",
			command: commands.MixerAnchor{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				X:        ptr(float32(0.0)),
				Y:        ptr(float32(0.1)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 ANCHOR 0 0.1 50 easeinsine",
		},
		{
			name: "Get CROP",
			command: commands.MixerCrop{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 CROP",
		},
		{
			name: "Set CROP",
			command: commands.MixerCrop{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				LeftEdge:   ptr(float32(0.1)),
				TopEdge:    ptr(float32(0.2)),
				RightEdge:  ptr(float32(0.3)),
				BottomEdge: ptr(float32(0.4)),
			},
			expected: "MIXER 1-2 CROP 0.1 0.2 0.3 0.4",
		},
		{
			name: "Set CROP with fade",
			command: commands.MixerCrop{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				LeftEdge:   ptr(float32(0.1)),
				TopEdge:    ptr(float32(0.2)),
				RightEdge:  ptr(float32(0.3)),
				BottomEdge: ptr(float32(0.4)),
				Duration:   ptr(50),
				Tween:      ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 CROP 0.1 0.2 0.3 0.4 50 easeinsine",
		},
		{
			name: "Get ROTATION",
			command: commands.MixerRotation{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 ROTATION",
		},
		{
			name: "Set ROTATION",
			command: commands.MixerRotation{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Angle: ptr(float32(120.2)),
			},
			expected: "MIXER 1-2 ROTATION 120.2",
		},
		{
			name: "Set ROTATION with fade",
			command: commands.MixerRotation{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Angle:    ptr(float32(120.2)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 ROTATION 120.2 50 easeinsine",
		},
		{
			name: "Get PERSPECTIVE",
			command: commands.MixerPerspective{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 PERSPECTIVE",
		},
		{
			name: "Set PERSPECTIVE",
			command: commands.MixerPerspective{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				TopLeftX:     ptr(float32(0.1)),
				TopLeftY:     ptr(float32(0.2)),
				TopRightX:    ptr(float32(0.3)),
				TopRightY:    ptr(float32(0.4)),
				BottomRightX: ptr(float32(0.5)),
				BottomRightY: ptr(float32(0.6)),
				BottomLeftX:  ptr(float32(0.7)),
				BottomLeftY:  ptr(float32(0.8)),
			},
			expected: "MIXER 1-2 PERSPECTIVE 0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8",
		},
		{
			name: "Set PERSPECTIVE with fade",
			command: commands.MixerPerspective{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				TopLeftX:     ptr(float32(0.1)),
				TopLeftY:     ptr(float32(0.2)),
				TopRightX:    ptr(float32(0.3)),
				TopRightY:    ptr(float32(0.4)),
				BottomRightX: ptr(float32(0.5)),
				BottomRightY: ptr(float32(0.6)),
				BottomLeftX:  ptr(float32(0.7)),
				BottomLeftY:  ptr(float32(0.8)),
				Duration:     ptr(50),
				Tween:        ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 PERSPECTIVE 0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 50 easeinsine",
		},
		{
			name: "Get MIPMAP",
			command: commands.MixerMipMap{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 MIPMAP",
		},
		{
			name: "Set MIPMAP true",
			command: commands.MixerMipMap{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Enable: ptr(true),
			},
			expected: "MIXER 1-2 MIPMAP 1",
		},
		{
			name: "Set MIPMAP false",
			command: commands.MixerMipMap{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Enable: ptr(false),
			},
			expected: "MIXER 1-2 MIPMAP 0",
		},
		{
			name: "Get VOLUME",
			command: commands.MixerVolume{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
			},
			expected: "MIXER 1-2 VOLUME",
		},
		{
			name: "Set VOLUME",
			command: commands.MixerVolume{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Volume: ptr(float32(0.1)),
			},
			expected: "MIXER 1-2 VOLUME 0.1",
		},
		{
			name: "Set VOLUME with fade",
			command: commands.MixerVolume{
				MixerCommand: commands.MixerCommand{
					VideoChannel: 1,
					Layer:        ptr(2),
				},
				Volume:   ptr(float32(0.1)),
				Duration: ptr(50),
				Tween:    ptr(types.TweenTypeEaseInSine),
			},
			expected: "MIXER 1-2 VOLUME 0.1 50 easeinsine",
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
