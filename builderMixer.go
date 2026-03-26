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
	layer        *int
	fade         *types.Fade // fade to apply to the next operation
}

// Mixer creates a new MixerBuilder for the specified video channel and layer.
func (c *Client) Mixer(videoChannel int, layer *int) *MixerBuilder {
	return &MixerBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// sendCommand abstracts sending a command that does not expect a response value.
func (b *MixerBuilder) sendCommand(cmd interface{ String() string }) error {
	_, err := b.client.Send(cmd)
	return err
}

// Fade sets a fade transition to be applied to the next setter operation.
// The fade is automatically cleared after being used, so it only applies to one operation.
// Returns the builder for method chaining.
func (b *MixerBuilder) Fade(fade *types.Fade) *MixerBuilder {
	b.fade = fade
	return b
}

// baseMixerCommand returns a base mixer command with channel and layer set.
func (b *MixerBuilder) baseMixerCommand() (commands.MixerCommand, error) {
	if b.layer == nil {
		return commands.MixerCommand{}, ErrLayerNotSet
	}
	return commands.MixerCommand{
		VideoChannel: b.videoChannel,
		Layer:        *b.layer,
	}, nil
}

// getIntValue retrieves an int value using the provided command.
func (b *MixerBuilder) getIntValue(cmd interface{ String() string }) (int, error) {
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}
	return responses.IntFromResponse(resp)
}

// getFloat32Value retrieves a float32 value using the provided command.
func (b *MixerBuilder) getFloat32Value(cmd interface{ String() string }) (float32, error) {
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}
	return responses.FloatFromResponse(resp)
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
	mc, err := b.baseMixerCommand()
	if err != nil {
		return false, err
	}
	cmd := commands.MixerCommandKeyer{
		MixerCommand: mc,
	}
	return b.getBoolValue(cmd)
}

// SetKeyer enables or disables the keyer effect.
// When enabled, layer n+1's alpha is replaced with the R (red) channel of layer n,
// and the RGB channels of layer n are hidden.
// If show is true, the specified layer will not be rendered; instead it will be used
// as the key for the layer above.
func (b *MixerBuilder) SetKeyer(show bool) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandKeyer{
		MixerCommand: mc,
		Show:         &show,
	}
	return b.setBoolValue(cmd)
}

// ChromaBuilder provides a fluent interface for chroma key operations.
type ChromaBuilder struct {
	mixer  *MixerBuilder
	params responses.MixerChroma
}

// GetChroma retrieves the current chroma key settings for the layer.
func (b *MixerBuilder) GetChroma() (responses.MixerChroma, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return responses.MixerChroma{}, err
	}
	cmd := commands.MixerCommandChroma{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerChroma{}, err
	}
	return responses.MixerChromaFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetChroma configures chroma key parameters and returns a ChromaBuilder.
//
// Use .Enable() or .Disable() to apply the settings.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetChroma(params responses.MixerChroma) *ChromaBuilder {
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

func (b *MixerBuilder) applyChroma(enable bool, params *responses.MixerChroma) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandChroma{
		MixerCommand: mc,
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
	return b.sendCommand(cmd)
}

// GetBlendMode retrieves the current blend mode for the layer.
func (b *MixerBuilder) GetBlendMode() (types.BlendMode, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return "", err
	}
	cmd := commands.MixerCommandBlend{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return "", err
	}
	return responses.BlendModeFromResponse(resp)
}

// SetBlendMode sets the blend mode for the layer.
func (b *MixerBuilder) SetBlendMode(mode types.BlendMode) error {
	if err := mode.Validate(); err != nil {
		return err
	}
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}

	cmd := commands.MixerCommandBlend{
		MixerCommand: mc,
		BlendMode:    &mode,
	}
	return b.setBoolValue(cmd)
}

// GetInvert retrieves the current invert state for the layer.
func (b *MixerBuilder) GetInvert() (bool, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return false, err
	}
	cmd := commands.MixerCommandInvert{
		MixerCommand: mc,
	}
	return b.getBoolValue(cmd)
}

// SetInvert enables or disables color inversion for the layer.
func (b *MixerBuilder) SetInvert(state bool) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandInvert{
		MixerCommand: mc,
		Invert:       &state,
	}
	return b.setBoolValue(cmd)
}

// GetOpacity retrieves the current opacity value for the layer (0.0 to 1.0).
func (b *MixerBuilder) GetOpacity() (float32, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return 0, err
	}
	return b.getFloat32Value(commands.MixerCommandOpacity{
		MixerCommand: mc,
	})
}

// SetOpacity sets the opacity value for the layer.
//
// opacity: float32 - The new opacity value, 0.0 = completely transparent, 1.0 = fully opaque.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetOpacity(opacity float32) error {
	if err := inRangeFloat("opacity", opacity, 0.0, 1.0); err != nil {
		return err
	}
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandOpacity{
		MixerCommand: mc,
		Opacity:      &opacity,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetBrightness retrieves the current brightness value for the layer.
func (b *MixerBuilder) GetBrightness() (float32, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return 0, err
	}
	return b.getFloat32Value(commands.MixerCommandBrightness{
		MixerCommand: mc,
	})
}

// SetBrightness sets the brightness value for the layer.
//
// brightness: float32 - The new brightness value, 0.5 = original brightness, 0.0 = completely dark, 1.0 = double brightness.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetBrightness(brightness float32) error {
	if err := inRangeFloat("brightness", brightness, 0.0, 1.0); err != nil {
		return err
	}
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandBrightness{
		MixerCommand: mc,
		Brightness:   &brightness,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetSaturation retrieves the current saturation value for the layer.
func (b *MixerBuilder) GetSaturation() (float32, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return 0, err
	}
	return b.getFloat32Value(commands.MixerCommandSaturation{
		MixerCommand: mc,
	})
}

// SetSaturation sets the saturation value for the layer.
//
// saturation: float32 - The new saturation value, 0.5 = original saturation, 0.0 = completely desaturated (gray), 1.0 = double saturation.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetSaturation(saturation float32) error {
	if err := inRangeFloat("saturation", saturation, 0.0, 1.0); err != nil {
		return err
	}
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandSaturation{
		MixerCommand: mc,
		Saturation:   &saturation,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetContrast retrieves the current contrast value for the layer.
func (b *MixerBuilder) GetContrast() (float32, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return 0, err
	}
	return b.getFloat32Value(commands.MixerCommandContrast{
		MixerCommand: mc,
	})
}

// SetContrast sets the contrast value for the layer.
//
// contrast: float32 - The new contrast value, 0.5 = original contrast, 0.0 = completely gray, 1.0 = double contrast.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetContrast(contrast float32) error {
	if err := inRangeFloat("contrast", contrast, 0.0, 1.0); err != nil {
		return err
	}
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandContrast{
		MixerCommand: mc,
		Contrast:     &contrast,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetLevels retrieves the current levels settings for the layer.
func (b *MixerBuilder) GetLevels() (types.MixerLevels, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return types.MixerLevels{}, err
	}
	cmd := commands.MixerCommandLevels{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return types.MixerLevels{}, err
	}
	return types.MixerInfoLevelsFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetLevels adjusts the input/output levels and gamma for the layer.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetLevels(params types.MixerLevels) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandLevels{
		MixerCommand: mc,
		MinInput:     &params.MinInput,
		MaxInput:     &params.MaxInput,
		Gamma:        &params.Gamma,
		MinOutput:    &params.MinOutput,
		MaxOutput:    &params.MaxOutput,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetFill retrieves the current fill (position and scale) settings for the layer.
func (b *MixerBuilder) GetFill() (responses.MixerFill, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return responses.MixerFill{}, err
	}
	cmd := commands.MixerCommandFill{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerFill{}, err
	}
	return responses.MixerFillFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetFill scales and positions the video stream on the specified layer.
// The positioning and scaling is performed around the anchor point set by MIXER ANCHOR.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetFill(params types.MixerParamsFill) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCommandFill{
		MixerCommand: mc,
		X:            &params.X,
		Y:            &params.Y,
		XScale:       &params.XScale,
		YScale:       &params.YScale,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetClip retrieves the current clip (position and scale) settings for the layer.
func (b *MixerBuilder) GetClip() (responses.MixerClip, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return responses.MixerClip{}, err
	}
	cmd := commands.MixerClip{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerClip{}, err
	}
	return responses.MixerClipFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetClip defines the rectangular viewport where a layer is rendered thru on the screen without being affected by MIXER FILL, MIXER ROTATION and MIXER PERSPECTIVE.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetClip(params responses.MixerClip) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerClip{
		MixerCommand: mc,
		X:            &params.X,
		Y:            &params.Y,
		Width:        &params.Width,
		Height:       &params.Height,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetAnchor retrieves the current anchor point settings for the layer.
func (b *MixerBuilder) GetAnchor() (responses.MixerAnchor, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return responses.MixerAnchor{}, err
	}
	cmd := commands.MixerAnchor{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerAnchor{}, err
	}
	return responses.MixerAnchorFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetAnchor changes the anchor point of the specified layer.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetAnchor(params responses.MixerAnchor) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerAnchor{
		MixerCommand: mc,
		X:            &params.X,
		Y:            &params.Y,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetCrop retrieves the current crop settings for the layer.
func (b *MixerBuilder) GetCrop() (responses.MixerCrop, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return responses.MixerCrop{}, err
	}
	cmd := commands.MixerCrop{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerCrop{}, err
	}
	return responses.MixerCropFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetCrop sets the crop values for the layer.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetCrop(params types.MixerCrop) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerCrop{
		MixerCommand: mc,
	}
	cmd.LeftEdge = &params.LeftEdge
	cmd.TopEdge = &params.TopEdge
	cmd.RightEdge = &params.RightEdge
	cmd.BottomEdge = &params.BottomEdge

	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetRotation retrieves the current rotation angle for the layer in degrees.
//
// 0 = no rotation, 90 = 90 degrees clockwise, -90 = 90 degrees counterclockwise. Higher and lower values allowed.
func (b *MixerBuilder) GetRotation() (float32, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return 0, err
	}
	cmd := commands.MixerRotation{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}
	return responses.FloatFromResponse(resp)
}

// SetRotation sets the rotation angle for the layer in degrees.
//
// rotation: float32 - The new rotation angle, 0 = no rotation, 90 = 90 degrees clockwise, -90 = 90 degrees counterclockwise. Higher and lower values allowed.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetRotation(rotation float32) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerRotation{
		MixerCommand: mc,
		Angle:        &rotation,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetPerspective retrieves the current perspective settings for the layer.
func (b *MixerBuilder) GetPerspective() (responses.MixerPerspective, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return responses.MixerPerspective{}, err
	}
	cmd := commands.MixerPerspective{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.MixerPerspective{}, err
	}
	return responses.MixerPerspectiveFromResponse(strings.Split(strings.Join(resp, ""), " "))
}

// SetPerspective sets the perspective transformation for the layer.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetPerspective(params types.MixerPerspective) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerPerspective{
		MixerCommand: mc,
	}
	cmd.TopLeftX = &params.TopLeftX
	cmd.TopLeftY = &params.TopLeftY
	cmd.TopRightX = &params.TopRightX
	cmd.TopRightY = &params.TopRightY
	cmd.BottomLeftX = &params.BottomLeftX
	cmd.BottomLeftY = &params.BottomLeftY
	cmd.BottomRightX = &params.BottomRightX
	cmd.BottomRightY = &params.BottomRightY
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetMipMap retrieves the current mipmap state for the layer.
func (b *MixerBuilder) GetMipMap() (bool, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return false, err
	}
	cmd := commands.MixerMipMap{
		MixerCommand: mc,
	}
	return b.getBoolValue(cmd)
}

// SetMipMap enables or disables mipmapping for the layer.
func (b *MixerBuilder) SetMipMap(enable bool) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerMipMap{
		MixerCommand: mc,
		Enable:       &enable,
	}
	return b.setBoolValue(cmd)
}

// GetVolume retrieves the current audio volume for the layer.
//
// 1.0 = original volume, 0.5 = half volume, 2.0 = double volume. Higher and lower values allowed.
func (b *MixerBuilder) GetVolume() (float32, error) {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return 0, err
	}
	cmd := commands.MixerVolume{
		MixerCommand: mc,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return 0, err
	}
	return responses.FloatFromResponse(resp)
}

// SetVolume sets the audio volume for the layer.
//
// volume: float32 - The new volume, 1.0 = original volume, 0.5 = half volume, 2.0 = double volume. Higher and lower values allowed.
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetVolume(volume float32) error {
	mc, err := b.baseMixerCommand()
	if err != nil {
		return err
	}
	cmd := commands.MixerVolume{
		MixerCommand: mc,
		Volume:       &volume,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// GetMasterVolume retrieves the current master audio volume for the video channel.
//
// 1.0 = original volume, 0.5 = half volume, 2.0 = double volume. Higher and lower values allowed.
func (b *MixerBuilder) GetMasterVolume() (float32, error) {
	if b.layer != nil {
		return 0, ErrLayerSet
	}
	cmd := commands.MixerMasterVolume{
		VideoChannel: b.videoChannel,
	}
	return b.getFloat32Value(cmd)
}

// SetMasterVolume sets the master audio volume for the video channel.
//
// volume: float32 - The new master volume, 1.0 = original volume, 0.5 = half volume, 2.0 = double volume. Higher and lower values allowed.
func (b *MixerBuilder) SetMasterVolume(volume float32) error {
	if b.layer != nil {
		return ErrLayerSet
	}
	cmd := commands.MixerMasterVolume{
		VideoChannel: b.videoChannel,
		Volume:       &volume,
	}
	return b.sendCommand(cmd)
}

// GetStraightAlphaOutput retrieves the current straight alpha output state for the channel.
func (b *MixerBuilder) GetStraightAlphaOutput() (bool, error) {
	if b.layer != nil {
		return false, ErrLayerSet
	}
	cmd := commands.MixerStraightAlphaOutput{
		VideoChannel: b.videoChannel,
	}
	return b.getBoolValue(cmd)
}

// SetStraightAlphaOutput enables or disables straight alpha output for the channel.
//
// This only works per video channel, not per layer.
func (b *MixerBuilder) SetStraightAlphaOutput(enable bool) error {
	if b.layer != nil {
		return ErrLayerSet
	}
	cmd := commands.MixerStraightAlphaOutput{
		VideoChannel: b.videoChannel,
		Enable:       &enable,
	}
	return b.setBoolValue(cmd)
}

// GetGrid retrieves the current grid settings for the video channel.
func (b *MixerBuilder) GetGrid() (int, error) {
	if b.layer != nil {
		return 0, ErrLayerSet
	}
	cmd := commands.MixerGrid{
		VideoChannel: b.videoChannel,
	}
	return b.getIntValue(cmd)
}

// SetGrid creates a grid of video layer in ascending order of the layer index,
// i.e. if resolution equals 2 then a 2x2 grid of layers will be created starting from layer 1.
//
// resolution: int - The number of cells in the grid. (e.g: 2 = 2x2 grid, 3 = 3x3 grid, etc.)
//
// Use Fade() before calling this method to apply a smooth transition.
func (b *MixerBuilder) SetGrid(resolution int) error {
	if err := inRangeInt("resolution", resolution, 1, 9999); err != nil {
		return err
	}
	if b.layer != nil {
		return ErrLayerSet
	}
	cmd := commands.MixerGrid{
		VideoChannel: b.videoChannel,
		Resolution:   &resolution,
	}
	b.applyFade(func(d *int) { cmd.Duration = d }, func(t *types.TweenType) { cmd.Tween = t })
	return b.sendCommand(cmd)
}

// Commit all deferred mixer transforms on the specified channel. This ensures that all animations start at the same exact frame.
func (b *MixerBuilder) Commit() error {
	cmd := commands.MixerCommit{
		VideoChannel: b.videoChannel,
	}
	return b.sendCommand(cmd)
}

// Clear clears all transformations on a channel or layer
func (b *MixerBuilder) Clear() error {
	cmd := commands.MixerClear{
		VideoChannel: b.videoChannel,
		Layer:        b.layer,
	}
	return b.sendCommand(cmd)
}
