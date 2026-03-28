package commands

import (
	"github.com/overlayfox/casparcg-amcp-go/types"
)

// QueryCINF returns information about a media file.
type QueryCINF struct {
	Filename string
}

func (c QueryCINF) String() string {
	return "CINF " + quote(c.Filename)
}

// QueryCLS lists media files in the media folder.
// Use the command INFO PATHS to get the path to the media folder.
type QueryCLS struct {
	Directory *string
}

func (c QueryCLS) String() string {
	cmd := "CLS"
	return appendQuotedString(cmd, c.Directory)
}

// QueryFLS lists all fonts in the fonts folder.
// Use the command INFO PATHS to get the path to the fonts folder.
type QueryFLS struct{}

func (c QueryFLS) String() string {
	return "FLS"
}

// QueryTLS lists template files in the templates folder.
// Use the command INFO PATHS to get the path to the templates folder.
type QueryTLS struct {
	Directory *string
}

func (c QueryTLS) String() string {
	cmd := "TLS"
	return appendString(cmd, c.Directory)
}

// QueryVersion returns the version of specified component.
type QueryVersion struct {
	Component types.VersionInfo
}

func (c QueryVersion) String() string {
	if c.Component != "" {
		return "VERSION " + string(c.Component)
	}
	return "VERSION"
}

// QueryInfo retrieves a list of available channels.
type QueryInfo struct {
	Component types.InfoComponent
}

func (c QueryInfo) String() string {
	if c.Component != "" {
		return "INFO " + string(c.Component)
	}
	return "INFO"
}

// QueryInfoTemplate gets information about the specified template.
type QueryInfoTemplate struct {
	Template string
}

func (c QueryInfoTemplate) String() string {
	return "INFO TEMPLATE " + quote(c.Template)
}

type QueryDiag struct{}

func (c QueryDiag) String() string {
	return "DIAG"
}

type QueryGLInfo struct{}

func (c QueryGLInfo) String() string {
	return "GL INFO"
}

type QueryGLGC struct{}

func (c QueryGLGC) String() string {
	return "GL GC"
}

type QueryHelp struct {
	Command *string
}

func (c QueryHelp) String() string {
	cmd := "HELP"
	return appendString(cmd, c.Command)
}

type QueryHelpProducer struct {
	Producer *string
}

func (c QueryHelpProducer) String() string {
	cmd := "HELP PRODUCER"
	return appendString(cmd, c.Producer)
}

type QueryHelpConsumer struct {
	Consumer *string
}

func (c QueryHelpConsumer) String() string {
	cmd := "HELP CONSUMER"
	return appendString(cmd, c.Consumer)
}
