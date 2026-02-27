package types

import "fmt"

type FrameRate struct {
	Num int `xml:"num"`
	Den int `xml:"den"`
}

func (f FrameRate) Float() float64 {
	if f.Den == 0 {
		return 0
	}
	return float64(f.Num) / float64(f.Den)
}

func StringToFrameRate(s string) (FrameRate, error) {
	var frameRate FrameRate
	_, err := fmt.Sscanf(s, "%d/%d", &frameRate.Num, &frameRate.Den)
	if err != nil {
		return FrameRate{}, err
	}
	return frameRate, nil
}

type MediaTypes string

const (
	CLSTypeStill MediaTypes = "STILL"
	CLSTypeMovie MediaTypes = "MOVIE"
	CLSTypeAudio MediaTypes = "AUDIO"
)

type LogLevel string

const (
	LogLevelTrace LogLevel = "trace"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
)

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

type AspectRatio string

const (
	AspectRatioDefault AspectRatio = "default"
	AspectRatio43      AspectRatio = "4:3"
	AspectRatio169     AspectRatio = "16:9"
)

type StretchMode string

const (
	StretchModeNone          StretchMode = "none"
	StretchModeFill          StretchMode = "fill"
	StretchModeUniform       StretchMode = "uniform"
	StretchModeUniformToFill StretchMode = "uniform_to_fill"
)

type ColourSpace string

const (
	ColourSpaceRGB              ColourSpace = "rgb"
	ColourSpaceDataVideoFull    ColourSpace = "datavideo-full"
	ColourSpaceDataVideoLimited ColourSpace = "datavideo-limited"
)

type AudioChannelLayout string

const (
	AudioChannelLayoutMono   AudioChannelLayout = "mono"
	AudioChannelLayoutStereo AudioChannelLayout = "stereo"
	AudioChannelLayoutMatrix AudioChannelLayout = "matrix"
)

type DecklinkLatency string

const (
	DecklinkLatencyLow    DecklinkLatency = "low"
	DecklinkLatencyMedium DecklinkLatency = "medium"
	DecklinkLatencyHigh   DecklinkLatency = "high"
)

type DecklinkKeyer string

const (
	DecklinkKeyerExternal               DecklinkKeyer = "external"
	DecklinkKeyerExternalSeparateDevice DecklinkKeyer = "external-separate-device"
	DecklinkKeyerInternal               DecklinkKeyer = "internal"
	DecklinkKeyerDefault                DecklinkKeyer = "default"
)
