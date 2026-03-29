package responses

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type CINF struct {
	Filename     string           `xml:"filename"`
	Type         types.MediaTypes `xml:"type"`
	FileSize     int64            `xml:"filesize"`
	LastModified time.Time        `xml:"lastmodified"`
	FrameCount   int              `xml:"framecount"`
	FrameRate    types.FrameRate  `xml:"framerate"`
}

type QueryChannelInfo struct {
	ChannelIndex int
	VideoMode    types.VideoMode
	Status       string
}

func PartsToQueryChannelInfo(parts []string) (QueryChannelInfo, error) {
	if len(parts) != 3 {
		return QueryChannelInfo{}, fmt.Errorf("unexpected format for channel info: %s", strings.Join(parts, " "))
	}

	videoChannel, err := strconv.Atoi(parts[0])
	if err != nil {
		return QueryChannelInfo{}, fmt.Errorf("invalid video channel in channel info: %s", parts[0])
	}

	channelInfo := QueryChannelInfo{
		ChannelIndex: videoChannel,
		VideoMode:    types.VideoMode(parts[1]),
		Status:       parts[2],
	}
	return channelInfo, nil
}

func ResponseToQueryChannelInfo(response []string) ([]QueryChannelInfo, error) {
	channelInfo := make([]QueryChannelInfo, 0, len(response))
	for _, line := range response {
		parts := strings.Split(line, " ")
		info, err := PartsToQueryChannelInfo(parts)
		if err != nil {
			return nil, err
		}
		channelInfo = append(channelInfo, info)
	}
	return channelInfo, nil
}

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

	q.VideoMode = raw.Format
	q.Mixer.Audio.Volume = raw.Mixer.Audio.Volume
	q.Output = raw.Output

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

// GLInfo is the root struct representing the <gl> tag.
type GLInfo struct {
	XMLName xml.Name `xml:"gl"`
	Details Details  `xml:"details"`
	Summary Summary  `xml:"summary"`
}

// Details represents the <details> section.
// Using string types here since the tags in your example are empty/self-closing.
type Details struct {
	PooledDeviceBuffers string `xml:"pooled_device_buffers"`
	PooledHostBuffers   string `xml:"pooled_host_buffers"`
}

// Summary represents the <summary> section, containing the aggregated stats.
type Summary struct {
	PooledDeviceBuffers DeviceBuffersStats `xml:"pooled_device_buffers"`
	PooledHostBuffers   HostBuffersStats   `xml:"pooled_host_buffers"`
	AllHostBuffers      HostBuffersStats   `xml:"all_host_buffers"`
}

// DeviceBuffersStats represents the structure for device buffer metrics.
type DeviceBuffersStats struct {
	TotalCount int64 `xml:"total_count"`
	TotalSize  int64 `xml:"total_size"`
}

// HostBuffersStats represents the structure for both pooled and all host buffer metrics.
type HostBuffersStats struct {
	TotalReadCount  int64 `xml:"total_read_count"`
	TotalWriteCount int64 `xml:"total_write_count"`
	TotalReadSize   int64 `xml:"total_read_size"`
	TotalWriteSize  int64 `xml:"total_write_size"`
}
