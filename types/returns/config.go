package returns

import (
	"encoding/xml"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type CasparConfig struct {
	XMLName xml.Name `xml:"configuration"`

	Paths    Paths `xml:"paths"`
	Channels struct {
		Channel []Channel `xml:"channel"`
	} `xml:"channels"`
	Controllers Controller `xml:"controllers"`

	LockClearPhrase *string             `xml:"lock-clear-phrase,omitempty"`
	AMCP            *AMCP               `xml:"amcp,omitempty"`
	LogLevel        *types.AMCPLogLevel `xml:"log-level,omitempty"`

	TemplateHosts *struct {
		TemplateHost *[]TemplateHost `xml:"template-host,omitempty"`
	} `xml:"template-hosts,omitempty"`

	Flash *struct {
		BufferDepth *string `xml:"buffer-depth,omitempty"`
	} `xml:"flash,omitempty"`

	HTML *HTML `xml:"html,omitempty"`

	NDI *struct {
		AutoLoad *bool `xml:"auto-load,omitempty"`
	} `xml:"ndi,omitempty"`
}

type Paths struct {
	Media string `xml:"media-path"`
	Log   struct {
		Value   string `xml:",chardata"`
		Disable bool   `xml:"disable,attr"`
	} `xml:"log-path"`
	Data     string `xml:"data-path"`
	Template string `xml:"template-path"`
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

type Channel struct {
	VideoMode VideoMode `xml:"video-mode"`
	Consumers struct {
		Decklink    *[]ConsumerDecklink    `xml:"decklink,omitempty"`
		SystemAudio *[]ConsumerSystemAudio `xml:"system-audio,omitempty"`
		Screen      *[]ConsumerScreen      `xml:"screen,omitempty"`
		NewTekIVP   *[]struct{}            `xml:"newtek-ivp,omitempty"`
		NDI         *[]ConsumerNDI         `xml:"ndi,omitempty"`
		FFMPEG      *[]ConsumerFFMPEG      `xml:"ffmpeg,omitempty"`
	} `xml:"consumers"`
}

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

type ConsumerDecklink struct {
	Device        int             `xml:"device"`
	KeyDevice     int             `xml:"key-device"`
	EmbeddedAudio bool            `xml:"embedded-audio"`
	Latency       DecklinkLatency `xml:"latency"`
	Keyer         DecklinkKeyer   `xml:"keyer"`
	KeyOnly       bool            `xml:"key-only"`
	BufferDepth   int             `xml:"buffer-depth"`
}

type ChannelLayout string

const (
	ChannelLayoutMono   ChannelLayout = "mono"
	ChannelLayoutStereo ChannelLayout = "stereo"
	ChannelLayoutMatrix ChannelLayout = "matrix"
)

type ConsumerSystemAudio struct {
	ChannelLayout ChannelLayout `xml:"channel-layout"`
	Latency       int           `xml:"latency"`
}

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

type ConsumerScreen struct {
	Device      int         `xml:"device"`
	AspectRatio AspectRatio `xml:"aspect-ratio"`
	StretchMode StretchMode `xml:"stretch-mode"`
	Windowed    bool        `xml:"windowed"`
	KeyOnly     bool        `xml:"key-only"`
	VSync       bool        `xml:"vsync"`
	Borderless  bool        `xml:"borderless"`
	Interactive bool        `xml:"interactive"`
	AlwaysOnTop bool        `xml:"always-on-top"`
	XOffset     int         `xml:"x-offset"`
	YOffset     int         `xml:"y-offset"`
	Width       int         `xml:"width"`
	Height      int         `xml:"height"`
	SbsKeyer    bool        `xml:"sbs-keyer"`
	ColourSpace ColourSpace `xml:"colour-space"`
}

type ConsumerNDI struct {
	Name        string `xml:"name"`
	AllowFields bool   `xml:"allow-fields"`
}

type ConsumerFFMPEG struct {
	Path string `xml:"path"`
	Args string `xml:"args"`
}

type Controller struct {
	TCP struct {
		Port     int    `xml:"port"`
		Protocol string `xml:"protocol"`
	} `xml:"tcp"`
}

type AMCP struct {
	MediaServer struct {
		Host string `xml:"host"`
		Port int    `xml:"port"`
	} `xml:"media-server"`
}

type TemplateHost struct {
	VideoMode string `xml:"video-mode"`
	Filename  string `xml:"filename"`
	Width     int    `xml:"width"`
	Height    int    `xml:"height"`
}

type HTML struct {
	RemoteDebuggingPort int  `xml:"remote-debugging-port"`
	EnableGPU           bool `xml:"enable-gpu"`
}
