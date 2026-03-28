package commands

import (
	"github.com/overlayfox/casparcg-amcp-go/types"
)

// QueryCommandCINF returns information about a media file.
type QueryCommandCINF struct {
	Filename string
}

func (c QueryCommandCINF) String() string {
	return "CINF " + quote(c.Filename)
}

// QueryCommandCLS lists media files in the media folder.
// Use the command INFO PATHS to get the path to the media folder.
type QueryCommandCLS struct {
	Directory *string
}

func (c QueryCommandCLS) String() string {
	cmd := "CLS"
	return appendString(cmd, c.Directory)
}

// QueryCommandFLS lists all fonts in the fonts folder.
// Use the command INFO PATHS to get the path to the fonts folder.
type QueryCommandFLS struct{}

func (c QueryCommandFLS) String() string {
	return "FLS"
}

// QueryCommandTLS lists template files in the templates folder.
// Use the command INFO PATHS to get the path to the templates folder.
type QueryCommandTLS struct {
	Directory *string
}

func (c QueryCommandTLS) String() string {
	cmd := "TLS"
	return appendString(cmd, c.Directory)
}

// QueryCommandVersion returns the version of specified component.
type QueryCommandVersion struct {
	Component types.VersionInfo
}

func (c QueryCommandVersion) String() string {
	if c.Component != "" {
		return "VERSION " + string(c.Component)
	}
	return "VERSION"
}

// QueryCommandInfo retrieves a list of available channels.
type QueryCommandInfo struct {
	Component types.InfoComponent
}

func (c QueryCommandInfo) String() string {
	if c.Component != "" {
		return "INFO " + string(c.Component)
	}
	return "INFO"
}

// QueryCommandInfoTemplate gets information about the specified template.
type QueryCommandInfoTemplate struct {
	Template string
}

func (c QueryCommandInfoTemplate) String() string {
	return "INFO TEMPLATE " + quote(c.Template)
}

type QueryCommandDiag struct{}

func (c QueryCommandDiag) String() string {
	return "DIAG"
}

type QueryCommandGLInfo struct{}

func (c QueryCommandGLInfo) String() string {
	return "GL INFO"
}

type QueryCommandGLGC struct{}

func (c QueryCommandGLGC) String() string {
	return "GL GC"
}

type QueryCommandHelp struct {
	Command *string
}

func (c QueryCommandHelp) String() string {
	cmd := "HELP"
	return appendString(cmd, c.Command)
}

type QueryCommandHelpProducer struct {
	Producer *string
}

func (c QueryCommandHelpProducer) String() string {
	cmd := "HELP PRODUCER"
	return appendString(cmd, c.Producer)
}

type QueryCommandHelpConsumer struct {
	Consumer *string
}

func (c QueryCommandHelpConsumer) String() string {
	cmd := "HELP CONSUMER"
	return appendString(cmd, c.Consumer)
}
