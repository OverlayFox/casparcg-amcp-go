package returns

import (
	"time"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type CLS struct {
	Filename     string           `xml:"filename"`
	Type         types.MediaTypes `xml:"type"`
	FileSize     int64            `xml:"filesize"`
	LastModified time.Time        `xml:"lastmodified"`
	FrameCount   int              `xml:"framecount"`
	FrameRate    types.FrameRate  `xml:"framerate"`
}
