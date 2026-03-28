package responses

import (
	"encoding/xml"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

// QueryChannelInfoVerbose is your target struct for JSON/Application use.
type QueryChannelInfoVerbose struct {
	VideoMode types.VideoMode `json:"VideoMode"`
	FrameRate types.FrameRate `json:"FrameRate"`
	Mixer     struct {
		Audio struct {
			Volume []int `json:"Volume"`
		} `json:"Audio"`
	} `json:"Mixer"`
	Output InfoChannelOutput `json:"Output"`
}

// UnmarshalXML is the entry point for the XML decoder.
func (q *QueryChannelInfoVerbose) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// This "proxy" struct mirrors the XML exactly and uses only EXPORTED fields
	type xmlChannel struct {
		Format     types.VideoMode `xml:"format"`
		FrameRates []int           `xml:"framerate"`
		Mixer      struct {
			Audio struct {
				Volume []int `xml:"volume"`
			} `xml:"audio"`
		} `xml:"mixer"`
		Output InfoChannelOutput `xml:"output"`
	}

	var raw xmlChannel
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	// Now map the raw XML data to your clean struct
	q.VideoMode = raw.Format
	q.Mixer.Audio.Volume = raw.Mixer.Audio.Volume
	q.Output = raw.Output

	// Handle the multiple <framerate> tags manually
	if len(raw.FrameRates) >= 2 {
		q.FrameRate.Num = raw.FrameRates[0]
		q.FrameRate.Den = raw.FrameRates[1]
	} else if len(raw.FrameRates) == 1 {
		q.FrameRate.Num = raw.FrameRates[0]
		q.FrameRate.Den = 1
	}

	return nil
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
