package returns

import (
	"encoding/xml"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type InfoChannel struct {
	VideoMode types.VideoMode `xml:"format"`
	FrameRate types.FrameRate `xml:"framerate"`
	Mixer     struct {
		Audio struct {
			Volume []int `xml:"volume"`
		} `xml:"audio"`
	} `xml:"mixer"`
	Output InfoChannelOutput `xml:"output"`
}

type InfoChannelOutput struct {
	Port struct {
		Consumers []InfoChannelConsumer `xml:",any"`
	} `xml:"port"`
}

type InfoChannelConsumer struct {
	XMLName  xml.Name `xml:""`
	Consumer string   `xml:"consumer"`
	Screen   *struct {
		AlwaysOnTop bool   `xml:"always_on_top"`
		Index       int    `xml:"index"`
		KeyOnly     bool   `xml:"key_only"`
		Name        string `xml:"name"`
	} `xml:"screen,omitempty"`
}
