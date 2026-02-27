package returns

import (
	"fmt"
	"time"
)

type CLSType string

const (
	CLSTypeStill CLSType = "STILL"
	CLSTypeMovie CLSType = "MOVIE"
	CLSTypeAudio CLSType = "AUDIO"
)

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

type CLS struct {
	Filename     string    `xml:"filename"`
	Type         CLSType   `xml:"type"`
	FileSize     int64     `xml:"filesize"`
	LastModified time.Time `xml:"lastmodified"`
	FrameCount   int       `xml:"framecount"`
	FrameRate    FrameRate `xml:"framerate"`
}
