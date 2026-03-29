package commands

import (
	"github.com/overlayfox/casparcg-amcp-go/types"
)

type QueryCINF struct {
	Filename string
}

func (c QueryCINF) String() string {
	return "CINF " + quote(c.Filename)
}

type QueryCLS struct {
	Directory *string
}

func (c QueryCLS) String() string {
	cmd := "CLS"
	return appendQuotedString(cmd, c.Directory)
}

type QueryFLS struct{}

func (c QueryFLS) String() string {
	return "FLS"
}

type QueryTLS struct {
	Directory *string
}

func (c QueryTLS) String() string {
	cmd := "TLS"
	return appendString(cmd, c.Directory)
}

type QueryVersion struct {
	Component types.VersionInfo
}

func (c QueryVersion) String() string {
	if c.Component != "" {
		return "VERSION " + string(c.Component)
	}
	return "VERSION"
}

type QueryInfo struct {
	Component types.InfoComponent
}

func (c QueryInfo) String() string {
	if c.Component != "" {
		return "INFO " + string(c.Component)
	}
	return "INFO"
}

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
