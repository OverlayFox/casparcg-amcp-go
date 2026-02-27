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

	LockClearPhrase *string         `xml:"lock-clear-phrase,omitempty"`
	AMCP            *AMCP           `xml:"amcp,omitempty"`
	LogLevel        *types.LogLevel `xml:"log-level,omitempty"`

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

type Channel struct {
	VideoMode types.VideoMode `xml:"video-mode"`
	Consumers struct {
		Decklink    *[]ConsumerDecklink    `xml:"decklink,omitempty"`
		SystemAudio *[]ConsumerSystemAudio `xml:"system-audio,omitempty"`
		Screen      *[]ConsumerScreen      `xml:"screen,omitempty"`
		NewTekIVP   *[]struct{}            `xml:"newtek-ivp,omitempty"`
		NDI         *[]ConsumerNDI         `xml:"ndi,omitempty"`
		FFMPEG      *[]ConsumerFFMPEG      `xml:"ffmpeg,omitempty"`
	} `xml:"consumers"`
}

type ConsumerDecklink struct {
	Device        int                   `xml:"device"`
	KeyDevice     int                   `xml:"key-device"`
	EmbeddedAudio bool                  `xml:"embedded-audio"`
	Latency       types.DecklinkLatency `xml:"latency"`
	Keyer         types.DecklinkKeyer   `xml:"keyer"`
	KeyOnly       bool                  `xml:"key-only"`
	BufferDepth   int                   `xml:"buffer-depth"`
}

type ConsumerSystemAudio struct {
	ChannelLayout types.AudioChannelLayout `xml:"channel-layout"`
	Latency       int                      `xml:"latency"`
}

type ConsumerScreen struct {
	Device      int               `xml:"device"`
	AspectRatio types.AspectRatio `xml:"aspect-ratio"`
	StretchMode types.StretchMode `xml:"stretch-mode"`
	Windowed    bool              `xml:"windowed"`
	KeyOnly     bool              `xml:"key-only"`
	VSync       bool              `xml:"vsync"`
	Borderless  bool              `xml:"borderless"`
	Interactive bool              `xml:"interactive"`
	AlwaysOnTop bool              `xml:"always-on-top"`
	XOffset     int               `xml:"x-offset"`
	YOffset     int               `xml:"y-offset"`
	Width       int               `xml:"width"`
	Height      int               `xml:"height"`
	SbsKeyer    bool              `xml:"sbs-keyer"`
	ColourSpace types.ColourSpace `xml:"colour-space"`
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
