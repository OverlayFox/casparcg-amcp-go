package types

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

//
// Some types are uppercase and some lowercase to match the expected input/output of the AMCP commands and responses.
//

type FrameRate struct {
	Num int `xml:"num"`
	Den int `xml:"den"`
}

func StringToFrameRate(s string) (FrameRate, error) {
	var frameRate FrameRate
	_, err := fmt.Sscanf(s, "%d/%d", &frameRate.Num, &frameRate.Den)
	if err != nil {
		return FrameRate{}, err
	}
	return frameRate, nil
}

// UnmarshalXML handles the repeating <framerate> tags manually.
func (f *FrameRate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var values []int
	for {
		var val int
		err := d.DecodeElement(&val, &start)
		if err != nil {
			return err
		}
		values = append(values, val)

		t, err := d.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		if nextStart, ok := t.(xml.StartElement); ok && nextStart.Name.Local == "framerate" {
			start = nextStart
			continue
		} else {
			break
		}
	}

	// Assign values to the struct
	if len(values) >= 1 {
		f.Num = values[0]
	}
	if len(values) >= 2 {
		f.Den = values[1]
	} else {
		f.Den = 1 // Default denominator
	}

	return nil
}

func (f *FrameRate) Float() float64 {
	if f.Den == 0 {
		return 0
	}
	return float64(f.Num) / float64(f.Den)
}

func (f *FrameRate) String() string {
	return fmt.Sprintf("%d/%d", f.Num, f.Den)
}

type MediaTypes string

const (
	CLSTypeStill MediaTypes = "STILL"
	CLSTypeMovie MediaTypes = "MOVIE"
	CLSTypeAudio MediaTypes = "AUDIO"
)

var validMediaTypes = map[MediaTypes]bool{
	CLSTypeStill: true,
	CLSTypeMovie: true,
	CLSTypeAudio: true,
}

func StringToMediaType(s string) (MediaTypes, error) {
	media := MediaTypes(strings.ToUpper(s))
	if ok := validMediaTypes[media]; ok {
		return media, nil
	}
	return "", fmt.Errorf("invalid media type: %s", s)
}

func (m MediaTypes) String() string {
	return string(m)
}

type LogLevel string

const (
	LogLevelTrace LogLevel = "TRACE"
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
	LogLevelFatal LogLevel = "FATAL"
)

var validLogLevels = map[LogLevel]bool{
	LogLevelTrace: true,
	LogLevelDebug: true,
	LogLevelInfo:  true,
	LogLevelWarn:  true,
	LogLevelError: true,
	LogLevelFatal: true,
}

func StringToLogLevel(s string) (LogLevel, error) {
	level := LogLevel(strings.ToUpper(s))
	if ok := validLogLevels[level]; !ok {
		return "", fmt.Errorf("invalid log level: %s", s)
	}
	return level, nil
}

func (l LogLevel) String() string {
	return string(l)
}

type VideoMode string

const (
	VideoModePAL          VideoMode = "PAL"
	VideoModeNTSC         VideoMode = "NTSC"
	VideoMode576p2500     VideoMode = "576p2500"
	VideoMode720p2398     VideoMode = "720p2398"
	VideoMode720p2400     VideoMode = "720p2400"
	VideoMode720p2500     VideoMode = "720p2500"
	VideoMode720p5000     VideoMode = "720p5000"
	VideoMode720p2997     VideoMode = "720p2997"
	VideoMode720p5994     VideoMode = "720p5994"
	VideoMode720p3000     VideoMode = "720p3000"
	VideoMode720p6000     VideoMode = "720p6000"
	VideoMode1080p2398    VideoMode = "1080p2398"
	VideoMode1080p2400    VideoMode = "1080p2400"
	VideoMode1080i5000    VideoMode = "1080i5000"
	VideoMode1080i5994    VideoMode = "1080i5994"
	VideoMode1080i6000    VideoMode = "1080i6000"
	VideoMode1080p2500    VideoMode = "1080p2500"
	VideoMode1080p2997    VideoMode = "1080p2997"
	VideoMode1080p3000    VideoMode = "1080p3000"
	VideoMode1080p5000    VideoMode = "1080p5000"
	VideoMode1080p5994    VideoMode = "1080p5994"
	VideoMode1080p6000    VideoMode = "1080p6000"
	VideoMode1556p2398    VideoMode = "1556p2398"
	VideoMode1556p2400    VideoMode = "1556p2400"
	VideoMode1556p2500    VideoMode = "1556p2500"
	VideoModeDCI1080p2398 VideoMode = "dci1080p2398"
	VideoModeDCI1080p2400 VideoMode = "dci1080p2400"
	VideoModeDCI1080p2500 VideoMode = "dci1080p2500"
	VideoMode2160p2398    VideoMode = "2160p2398"
	VideoMode2160p2400    VideoMode = "2160p2400"
	VideoMode2160p2500    VideoMode = "2160p2500"
	VideoMode2160p2997    VideoMode = "2160p2997"
	VideoMode2160p3000    VideoMode = "2160p3000"
	VideoMode2160p5000    VideoMode = "2160p5000"
	VideoMode2160p5994    VideoMode = "2160p5994"
	VideoMode2160p6000    VideoMode = "2160p6000"
	VideoModeDCI2160p2398 VideoMode = "dci2160p2398"
	VideoModeDCI2160p2400 VideoMode = "dci2160p2400"
	VideoModeDCI2160p2500 VideoMode = "dci2160p2500"
)

var validVideoModes = map[VideoMode]bool{
	VideoModePAL:          true,
	VideoModeNTSC:         true,
	VideoMode576p2500:     true,
	VideoMode720p2398:     true,
	VideoMode720p2400:     true,
	VideoMode720p2500:     true,
	VideoMode720p5000:     true,
	VideoMode720p2997:     true,
	VideoMode720p5994:     true,
	VideoMode720p3000:     true,
	VideoMode720p6000:     true,
	VideoMode1080p2398:    true,
	VideoMode1080p2400:    true,
	VideoMode1080i5000:    true,
	VideoMode1080i5994:    true,
	VideoMode1080i6000:    true,
	VideoMode1080p2500:    true,
	VideoMode1080p2997:    true,
	VideoMode1080p3000:    true,
	VideoMode1080p5000:    true,
	VideoMode1080p5994:    true,
	VideoMode1080p6000:    true,
	VideoMode1556p2398:    true,
	VideoMode1556p2400:    true,
	VideoMode1556p2500:    true,
	VideoModeDCI1080p2398: true,
	VideoModeDCI1080p2400: true,
	VideoModeDCI1080p2500: true,
	VideoMode2160p2398:    true,
	VideoMode2160p2400:    true,
	VideoMode2160p2500:    true,
	VideoMode2160p2997:    true,
	VideoMode2160p3000:    true,
	VideoMode2160p5000:    true,
	VideoMode2160p5994:    true,
	VideoMode2160p6000:    true,
	VideoModeDCI2160p2398: true,
	VideoModeDCI2160p2400: true,
	VideoModeDCI2160p2500: true,
}

func StringToVideoMode(s string) (VideoMode, error) {
	mode := VideoMode(strings.ToLower(s))
	if ok := validVideoModes[mode]; !ok {
		return "", fmt.Errorf("invalid video mode: %s", s)
	}
	return mode, nil
}

func (v VideoMode) String() string {
	return string(v)
}

type AspectRatio string

const (
	AspectRatioDefault AspectRatio = "default"
	AspectRatio43      AspectRatio = "4:3"
	AspectRatio169     AspectRatio = "16:9"
)

var validAspectRatios = map[AspectRatio]bool{
	AspectRatioDefault: true,
	AspectRatio43:      true,
	AspectRatio169:     true,
}

func StringToAspectRatio(s string) (AspectRatio, error) {
	ratio := AspectRatio(strings.ToLower(s))
	if ok := validAspectRatios[ratio]; !ok {
		return "", fmt.Errorf("invalid aspect ratio: %s", s)
	}
	return ratio, nil
}

func (a AspectRatio) String() string {
	return string(a)
}

type StretchMode string

const (
	StretchModeNone          StretchMode = "none"
	StretchModeFill          StretchMode = "fill"
	StretchModeUniform       StretchMode = "uniform"
	StretchModeUniformToFill StretchMode = "uniform_to_fill"
)

var validStretchModes = map[StretchMode]bool{
	StretchModeNone:          true,
	StretchModeFill:          true,
	StretchModeUniform:       true,
	StretchModeUniformToFill: true,
}

func StringToStretchMode(s string) (StretchMode, error) {
	mode := StretchMode(strings.ToLower(s))
	if ok := validStretchModes[mode]; !ok {
		return "", fmt.Errorf("invalid stretch mode: %s", s)
	}
	return mode, nil
}

func (s StretchMode) String() string {
	return string(s)
}

type ColourSpace string

const (
	ColourSpaceRGB              ColourSpace = "rgb"
	ColourSpaceDataVideoFull    ColourSpace = "datavideo-full"
	ColourSpaceDataVideoLimited ColourSpace = "datavideo-limited"
)

var validColourSpaces = map[ColourSpace]bool{
	ColourSpaceRGB:              true,
	ColourSpaceDataVideoFull:    true,
	ColourSpaceDataVideoLimited: true,
}

func StringToColourSpace(s string) (ColourSpace, error) {
	space := ColourSpace(strings.ToLower(s))
	if ok := validColourSpaces[space]; !ok {
		return "", fmt.Errorf("invalid colour space: %s", s)
	}
	return space, nil
}

func (c ColourSpace) String() string {
	return string(c)
}

type AudioChannelLayout string

const (
	AudioChannelLayoutMono   AudioChannelLayout = "mono"
	AudioChannelLayoutStereo AudioChannelLayout = "stereo"
	AudioChannelLayoutMatrix AudioChannelLayout = "matrix"
)

var validAudioChannelLayouts = map[AudioChannelLayout]bool{
	AudioChannelLayoutMono:   true,
	AudioChannelLayoutStereo: true,
	AudioChannelLayoutMatrix: true,
}

func StringToAudioChannelLayout(s string) (AudioChannelLayout, error) {
	layout := AudioChannelLayout(strings.ToLower(s))
	if ok := validAudioChannelLayouts[layout]; !ok {
		return "", fmt.Errorf("invalid audio channel layout: %s", s)
	}
	return layout, nil
}

func (a AudioChannelLayout) String() string {
	return string(a)
}

type DecklinkLatency string

const (
	DecklinkLatencyLow    DecklinkLatency = "low"
	DecklinkLatencyMedium DecklinkLatency = "medium"
	DecklinkLatencyHigh   DecklinkLatency = "high"
)

var validDecklinkLatencies = map[DecklinkLatency]bool{
	DecklinkLatencyLow:    true,
	DecklinkLatencyMedium: true,
	DecklinkLatencyHigh:   true,
}

func StringToDecklinkLatency(s string) (DecklinkLatency, error) {
	latency := DecklinkLatency(strings.ToLower(s))
	if ok := validDecklinkLatencies[latency]; !ok {
		return "", fmt.Errorf("invalid decklink latency: %s", s)
	}
	return latency, nil
}

func (d DecklinkLatency) String() string {
	return string(d)
}

type DecklinkKeyer string

const (
	DecklinkKeyerExternal               DecklinkKeyer = "external"
	DecklinkKeyerExternalSeparateDevice DecklinkKeyer = "external-separate-device"
	DecklinkKeyerInternal               DecklinkKeyer = "internal"
	DecklinkKeyerDefault                DecklinkKeyer = "default"
)

var validDecklinkKeyers = map[DecklinkKeyer]bool{
	DecklinkKeyerExternal:               true,
	DecklinkKeyerExternalSeparateDevice: true,
	DecklinkKeyerInternal:               true,
	DecklinkKeyerDefault:                true,
}

func StringToDecklinkKeyer(s string) (DecklinkKeyer, error) {
	keyer := DecklinkKeyer(strings.ToLower(s))
	if ok := validDecklinkKeyers[keyer]; !ok {
		return "", fmt.Errorf("invalid decklink keyer: %s", s)
	}
	return keyer, nil
}

func (d DecklinkKeyer) String() string {
	return string(d)
}

type TweenType string

const (
	TweenTypeLinear     TweenType = "linear"
	TweenTypeEaseInSine TweenType = "easeinsine"
)

var validTweenTypes = map[TweenType]bool{
	TweenTypeLinear:     true,
	TweenTypeEaseInSine: true,
}

func ParseTweenType(s string) (TweenType, error) {
	tweenType := TweenType(strings.ToLower(s))
	if ok := validTweenTypes[tweenType]; !ok {
		return "", fmt.Errorf("invalid tween type: %s", s)
	}
	return tweenType, nil
}

func (t TweenType) String() string {
	return string(t)
}

type BlendMode string

const (
	BlendModeNormal BlendMode = "NORMAL"
	BlendModeScreen BlendMode = "SCREEN"
)

var validBlendModes = map[BlendMode]bool{
	BlendModeNormal: true,
	BlendModeScreen: true,
}

func ParseBlendMode(s string) (BlendMode, error) {
	mode := BlendMode(strings.ToUpper(s))
	if ok := validBlendModes[mode]; !ok {
		return "", fmt.Errorf("invalid blend mode: %s", s)
	}
	return mode, nil
}

func (b BlendMode) String() string {
	return string(b)
}

func (b BlendMode) Validate() error {
	if ok := validBlendModes[b]; !ok {
		return fmt.Errorf("invalid blend mode: %s", b)
	}
	return nil
}

type VersionInfo string

const (
	VersionInfoServer       VersionInfo = "SERVER"
	VersionInfoFlash        VersionInfo = "FLASH"
	VersionInfoTemplateHost VersionInfo = "TEMPLATE_HOST"
	VersionInfoCEF          VersionInfo = "CEF"
)

var validVersionInfo = map[VersionInfo]bool{
	VersionInfoServer:       true,
	VersionInfoFlash:        true,
	VersionInfoTemplateHost: true,
	VersionInfoCEF:          true,
}

func StringToVersionInfo(s string) (VersionInfo, error) {
	info := VersionInfo(strings.ToUpper(s))
	if ok := validVersionInfo[info]; !ok {
		return "", fmt.Errorf("invalid version info type: %s", s)
	}
	return info, nil
}

func (v VersionInfo) String() string {
	return string(v)
}

type InfoComponent string

const (
	InfoComponentConfig  InfoComponent = "CONFIG"
	InfoComponentPaths   InfoComponent = "PATHS"
	InfoComponentSystem  InfoComponent = "SYSTEM"
	InfoComponentServer  InfoComponent = "SERVER"
	InfoComponentQueues  InfoComponent = "QUEUES"
	InfoComponentThreads InfoComponent = "THREADS"
)

var validInfoComponents = map[InfoComponent]bool{
	InfoComponentConfig:  true,
	InfoComponentPaths:   true,
	InfoComponentSystem:  true,
	InfoComponentServer:  true,
	InfoComponentQueues:  true,
	InfoComponentThreads: true,
}

func StringToInfoComponent(s string) (InfoComponent, error) {
	component := InfoComponent(strings.ToUpper(s))
	if ok := validInfoComponents[component]; !ok {
		return "", fmt.Errorf("invalid info component: %s", s)
	}
	return component, nil
}

func (i InfoComponent) String() string {
	return string(i)
}

type SetVariable string

const (
	SetVariableMode          SetVariable = "MODE"
	SetVariableChannelLayout SetVariable = "CHANNEL_LAYOUT"
)

var validSetVariables = map[SetVariable]bool{
	SetVariableMode:          true,
	SetVariableChannelLayout: true,
}

func StringToSetVariable(s string) (SetVariable, error) {
	variable := SetVariable(strings.ToUpper(s))
	if ok := validSetVariables[variable]; !ok {
		return "", fmt.Errorf("invalid set variable: %s", s)
	}
	return variable, nil
}

func (v SetVariable) String() string {
	return string(v)
}

type LockAction string

const (
	LockActionAcquire LockAction = "ACQUIRE"
	LockActionRelease LockAction = "RELEASE"
	LockActionClear   LockAction = "CLEAR"
)

var validLockActions = map[LockAction]bool{
	LockActionAcquire: true,
	LockActionRelease: true,
	LockActionClear:   true,
}

func StringToLockAction(s string) (LockAction, error) {
	action := LockAction(strings.ToUpper(s))
	if ok := validLockActions[action]; !ok {
		return "", fmt.Errorf("invalid lock action: %s", s)
	}
	return action, nil
}

func (a LockAction) String() string {
	return string(a)
}
