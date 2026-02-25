package types

import "encoding/xml"

type CasparConfig struct {
	XMLName xml.Name `xml:"configuration"`
	Paths   struct {
		Media string `xml:"media-path"`
		Log   struct {
			Value   string `xml:",chardata"`
			Disable bool   `xml:"disable,attr"`
		} `xml:"log-path"`
		Data     string `xml:"data-path"`
		Template string `xml:"template-path"`
	} `xml:"paths"`
	LockClearPhrase string `xml:"lock-clear-phrase"`
	Channels        []struct {
		VideoMode string `xml:"video-mode"`
		Consumers struct {
			Screen      *struct{} `xml:"screen"`
			SystemAudio *struct{} `xml:"system-audio"`
		} `xml:"consumers"`
	} `xml:"channels>channel"`
	Controllers struct {
		TCP struct {
			Port     int    `xml:"port"`
			Protocol string `xml:"protocol"`
		} `xml:"tcp"`
	} `xml:"controllers"`
	AMCP struct {
		MediaServer struct {
			Host string `xml:"host"`
			Port int    `xml:"port"`
		} `xml:"media-server"`
	} `xml:"amcp"`
}
