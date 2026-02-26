package types

import (
	"strconv"
)

type QueryCommandInterface interface {
	String() string
}

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
	if c.Directory != nil {
		return "CLS " + quote(*c.Directory)
	}
	return "CLS"
}

// QueryCommandFLS lists all fonts in the fonts folder.
// Use the command INFO PATHS to get the path to the fonts folder.
type QueryCommandFLS struct {
}

func (c QueryCommandFLS) String() string {
	return "FLS"
}

// QueryCommandTLS lists template files in the templates folder.
// Use the command INFO PATHS to get the path to the templates folder.
type QueryCommandTLS struct {
	Directory *string
}

func (c QueryCommandTLS) String() string {
	if c.Directory != nil {
		return "TLS " + quote(*c.Directory)
	}
	return "TLS"
}

// QueryCommandVersion returns the version of specified component.
type QueryCommandVersion struct {
	Component *string
}

func (c QueryCommandVersion) String() string {
	if c.Component != nil {
		return "VERSION " + *c.Component
	}
	return "VERSION"
}

type InfoComponent string

const (
	InfoComponentConfig  InfoComponent = "CONFIG"
	InfoComponentPaths   InfoComponent = "PATHS"
	InfoComponentSystem  InfoComponent = "SYSTEM"
	InfoComponentServer  InfoComponent = "SERVER"
	InfoComponentQueues  InfoComponent = "QUEUES"
	InfoComponentThreads InfoComponent = "THREADS"
)

// QueryCommandInfo retrieves a list of available channels
type QueryCommandInfo struct {
	Component InfoComponent
}

func (c QueryCommandInfo) String() string {
	if c.Component != "" {
		return "INFO " + string(c.Component)
	}
	return "INFO"
}

// QueryCommandInfoChannel get information about a channel or a specific layer on a channel.
// If layer is ommitted information about the whole channel is returned.
type QueryCommandInfoChannel struct {
	VideoChannel int
	Layer        *int
}

func (c QueryCommandInfoChannel) String() string {
	cmd := "INFO " + strconv.Itoa(c.VideoChannel)
	if c.Layer != nil {
		cmd += "-" + strconv.Itoa(*c.Layer)
	}
	return cmd
}

// QueryCommandInfoTemplate gets information about the specified template.
type QueryCommandInfoTemplate struct {
	Template string
}

func (c QueryCommandInfoTemplate) String() string {
	return "INFO TEMPLATE " + quote(c.Template)
}

type QueryCommandInfoDelay struct {
	VideoChannel int
	Layer        *int
}

func (c QueryCommandInfoDelay) String() string {
	cmd := "INFO " + strconv.Itoa(c.VideoChannel)
	if c.Layer != nil {
		cmd += "-" + strconv.Itoa(*c.Layer)
	}
	cmd += " DELAY"
	return cmd
}
