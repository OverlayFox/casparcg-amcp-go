package casparcg

import (
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
	"github.com/overlayfox/casparcg-amcp-go/types/responses"
)

// MixerBuilder provides a fluent interface for mixer operations on a specific video channel and layer.
type MixerBuilder struct {
	client       *Client
	videoChannel int
	layer        int
	fade         *types.Fade // fade to apply to the next operation
}

// Mixer creates a new MixerBuilder for the specified video channel and layer.
func (c *Client) Mixer(videoChannel, layer int) *MixerBuilder {
	return &MixerBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// Fade sets a fade transition to be applied to the next setter operation.
// The fade is automatically cleared after being used, so it only applies to one operation.
// Returns the builder for method chaining.
func (b *MixerBuilder) Fade(fade *types.Fade) *MixerBuilder {
	b.fade = fade
	return b
}

// baseMixerCommand returns a base mixer command with channel and layer set.
func (b *MixerBuilder) baseMixerCommand() commands.MixerCommand {
	return commands.MixerCommand{
		VideoChannel: b.videoChannel,
		Layer:        b.layer,
	}
}

// getFloat32Value retrieves a float32 value using the provided command.
func (b *MixerBuilder) getFloat32Value(cmd interface{ String() string }) (float32, error) {
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}
	return responses.FloatFromResponse(resp)
}

// setFloat32Value sets a float32 value using the provided command.
func (b *MixerBuilder) setFloat32Value(cmd interface{ String() string }) error {
	_, err := b.client.Send(cmd)
	return err
}

// getBoolValue retrieves a boolean value using the provided command.
func (b *MixerBuilder) getBoolValue(cmd interface{ String() string }) (bool, error) {
	resp, err := b.client.Send(cmd)
	if err != nil {
		return false, err
	}
	return responses.BoolFromResponse(resp)
}

// setBoolValue sets a boolean value using the provided command.
func (b *MixerBuilder) setBoolValue(cmd interface{ String() string }) error {
	_, err := b.client.Send(cmd)
	return err
}

// applyFade applies fade parameters to a command if fade is not nil.
// This is a helper to reduce repetition across multiple setter methods.
func (b *MixerBuilder) applyFade(setDuration func(*int), setTween func(*types.TweenType)) {
	if b.fade != nil {
		setDuration(&b.fade.Duration)
		setTween(&b.fade.Tween)
		b.fade = nil // Clear fade after applying it
	}
}

// GetKeyer retrieves the current keyer state for the layer.
func (b *MixerBuilder) GetKeyer() (bool, error) {
	cmd := commands.MixerCommandKeyer{
		MixerCommand: b.baseMixerCommand(),
	}
	return b.getBoolValue(cmd)
}

// SetKeyer enables or disables the keyer effect.
// When enabled, layer n+1's alpha is replaced with the R (red) channel of layer n,
// and the RGB channels of layer n are hidden.
// If show is true, the specified layer will not be rendered; instead it will be used
// as the key for the layer above.
func (b *MixerBuilder) SetKeyer(show bool) error {
	cmd := commands.MixerCommandKeyer{
		MixerCommand: b.baseMixerCommand(),
		Show:         &show,
	}
	return b.setBoolValue(cmd)
}

// ChromaBuilder provides a fluent interface for chroma key operations.
type ChromaBuilder struct {
	mixer  *MixerBuilder
	params responses.MixerInfoChroma
}

// GetChroma retrieves the current chroma key settings for the layer.
func (b *MixerBuilder) GetChroma() (responses.MixerInfoChroma, error) {
	cmd := commands.MixerCommandChroma{
		MixerCommand: b.baseMixerCommand(),
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerInfoChroma{}, err
	}
	return responses.MixerInfoChromaFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetChroma configures chroma key parameters and returns a ChromaBuilder.
// Use .Enable() or .Disable() to apply the settings.
// Use mixer.Fade() before calling SetChroma to apply a transition.
func (b *MixerBuilder) SetChroma(params responses.MixerInfoChroma) *ChromaBuilder {
	return &ChromaBuilder{
		mixer:  b,
		params: params,
	}
}

// Enable enables chroma keying with the configured parameters.
func (c *ChromaBuilder) Enable() error {
	return c.mixer.applyChroma(true, &c.params)
}

// Disable disables chroma keying.
func (c *ChromaBuilder) Disable() error {
	return c.mixer.applyChroma(false, nil)
}

// applyChroma is a helper method to enable/disable chroma with optional parameters.
func (b *MixerBuilder) applyChroma(enable bool, params *responses.MixerInfoChroma) error {
	cmd := commands.MixerCommandChroma{
		MixerCommand: b.baseMixerCommand(),
		Enable:       &enable,
	}

	if params != nil {
		cmd.TargetHue = &params.TargetHue
		cmd.HueWidth = &params.HueWidth
		cmd.MinSaturation = &params.MinSaturation
		cmd.MinBrightness = &params.MinBrightness
		cmd.Softness = &params.Softness
		cmd.SpillSuppress = &params.SpillSuppress
		cmd.SpillSuppressSaturation = &params.SpillSuppressSaturation
		cmd.ShowMask = &params.ShowMask
	}

	b.applyFade(func(d *int) { cmd.FadeDuration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}

// GetBlendMode retrieves the current blend mode for the layer.
func (b *MixerBuilder) GetBlendMode() (types.BlendMode, error) {
	cmd := commands.MixerCommandBlend{
		MixerCommand: b.baseMixerCommand(),
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return "", err
	}
	return responses.MixerBlendModeFromResponse(resp)
}

// SetBlendMode sets the blend mode for the layer.
func (b *MixerBuilder) SetBlendMode(mode types.BlendMode) error {
	cmd := commands.MixerCommandBlend{
		MixerCommand: b.baseMixerCommand(),
		BlendMode:    &mode,
	}
	return b.setBoolValue(cmd)
}

// GetInvert retrieves the current invert state for the layer.
func (b *MixerBuilder) GetInvert() (bool, error) {
	cmd := commands.MixerCommandInvert{
		MixerCommand: b.baseMixerCommand(),
	}
	return b.getBoolValue(cmd)
}

// SetInvert enables or disables color inversion for the layer.
func (b *MixerBuilder) SetInvert(state bool) error {
	cmd := commands.MixerCommandInvert{
		MixerCommand: b.baseMixerCommand(),
		Invert:       &state,
	}
	return b.setBoolValue(cmd)
}

// GetOpacity retrieves the current opacity value for the layer (0.0 to 1.0).
func (b *MixerBuilder) GetOpacity() (float32, error) {
	return b.getFloat32Value(commands.MixerCommandOpacity{
		MixerCommand: b.baseMixerCommand(),
	})
}

// SetOpacity sets the opacity value for the layer (0.0 to 1.0).
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetOpacity(opacity float32) error {
	cmd := commands.MixerCommandOpacity{
		MixerCommand: b.baseMixerCommand(),
		Opacity:      &opacity,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}

// GetBrightness retrieves the current brightness value for the layer.
func (b *MixerBuilder) GetBrightness() (float32, error) {
	return b.getFloat32Value(commands.MixerCommandBrightness{
		MixerCommand: b.baseMixerCommand(),
	})
}

// SetBrightness sets the brightness value for the layer.
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetBrightness(brightness float32) error {
	cmd := commands.MixerCommandBrightness{
		MixerCommand: b.baseMixerCommand(),
		Brightness:   &brightness,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}

// GetSaturation retrieves the current saturation value for the layer.
func (b *MixerBuilder) GetSaturation() (float32, error) {
	return b.getFloat32Value(commands.MixerCommandSaturation{
		MixerCommand: b.baseMixerCommand(),
	})
}

// SetSaturation sets the saturation value for the layer.
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetSaturation(saturation float32) error {
	cmd := commands.MixerCommandSaturation{
		MixerCommand: b.baseMixerCommand(),
		Saturation:   &saturation,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}

// GetContrast retrieves the current contrast value for the layer.
func (b *MixerBuilder) GetContrast() (float32, error) {
	return b.getFloat32Value(commands.MixerCommandContrast{
		MixerCommand: b.baseMixerCommand(),
	})
}

// SetContrast sets the contrast value for the layer.
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetContrast(contrast float32) error {
	cmd := commands.MixerCommandContrast{
		MixerCommand: b.baseMixerCommand(),
		Contrast:     &contrast,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}

// GetLevels retrieves the current levels settings for the layer.
func (b *MixerBuilder) GetLevels() (types.MixerInfoLevels, error) {
	cmd := commands.MixerCommandLevels{
		MixerCommand: b.baseMixerCommand(),
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return types.MixerInfoLevels{}, err
	}
	return types.MixerInfoLevelsFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetLevels adjusts the input/output levels and gamma for the layer.
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetLevels(params types.MixerInfoLevels) error {
	cmd := commands.MixerCommandLevels{
		MixerCommand: b.baseMixerCommand(),
		MinInput:     &params.MinInput,
		MaxInput:     &params.MaxInput,
		Gamma:        &params.Gamma,
		MinOutput:    &params.MinOutput,
		MaxOutput:    &params.MaxOutput,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}

// GetFill retrieves the current fill (position and scale) settings for the layer.
func (b *MixerBuilder) GetFill() (responses.MixerInfoFill, error) {
	cmd := commands.MixerCommandFill{
		MixerCommand: b.baseMixerCommand(),
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerInfoFill{}, err
	}
	return responses.MixerInfoFillFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetFill scales and positions the video stream on the specified layer.
// The positioning and scaling is performed around the anchor point set by MIXER ANCHOR.
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetFill(params types.MixerParamsFill) error {
	cmd := commands.MixerCommandFill{
		MixerCommand: b.baseMixerCommand(),
		X:            params.X,
		Y:            params.Y,
		XScale:       params.XScale,
		YScale:       params.YScale,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.setFloat32Value(cmd)
}
