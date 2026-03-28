package casparcg

import (
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
	"github.com/overlayfox/casparcg-amcp-go/types/responses"
)

type QueryBuilder struct {
	client *Client
}

// Query creates a new query command builder for the specified channel and layer.
func (c *Client) Query() *QueryBuilder {
	return &QueryBuilder{
		client: c,
	}
}

var reCINF = regexp.MustCompile(`^"?([^"]+)"?\s+(\S+)\s+(\d+)\s+(\d+)\s+(\d+)\s+([\d/]+)$`)

// CINF returns information about a media file.
func (b *QueryBuilder) CINF(filename string) (responses.CINF, error) {
	cmd := commands.QueryCommandCINF{
		Filename: filename,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.CINF{}, err
	}

	matches := reCINF.FindStringSubmatch(resp[0])
	cinf, err := matchesToCINF(matches)
	if err != nil {
		return responses.CINF{}, err
	}

	return cinf, nil
}

// CLS lists media files in the media folder.
//
// Use the command INFO PATHS to get the path to the media folder.
//
// If the optional sub_directory is specified only the media files in that sub directory will be returned.
func (b *QueryBuilder) CLS(subDirectory *string) ([]responses.CINF, error) {
	cmd := commands.QueryCommandCLS{
		Directory: subDirectory,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return nil, err
	}

	cls := []responses.CINF{}
	for _, file := range resp {
		matches := reCINF.FindStringSubmatch(file)
		cinf, err := matchesToCINF(matches)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cinf)
	}

	return cls, nil
}

// FLS lists all fonts
//
// TODO: implement DTO for FLS response
func (b *QueryBuilder) FLS() ([]string, error) {
	cmd := commands.QueryCommandFLS{}
	return b.client.Send(cmd)
}

// TLS lists template files.
//
// Use the command INFO PATHS to get the path to the templates folder.
//
// If the optional sub_directory is specified only the template files in that sub directory will be returned.
func (c *Client) TLS(directory *string) ([]string, error) {
	cmd := commands.QueryCommandTLS{
		Directory: directory,
	}
	return c.Send(cmd)
}

type QueryVersionCommand struct {
	QueryBuilder
}

// Version returns the version of specified component.
func (b *QueryBuilder) Version() *QueryVersionCommand {
	return &QueryVersionCommand{
		QueryBuilder: *b,
	}
}

func (b *QueryVersionCommand) Generic() (string, error) {
	return b.client.version("")
}

func (b *QueryVersionCommand) Server() (string, error) {
	return b.client.version(types.VersionInfoServer)
}

func (b *QueryVersionCommand) Flash() (string, error) {
	return b.client.version(types.VersionInfoFlash)
}

func (b *QueryVersionCommand) TemplateHost() (string, error) {
	return b.client.version(types.VersionInfoTemplateHost)
}

func (b *QueryVersionCommand) CEF() (string, error) {
	return b.client.version(types.VersionInfoCEF)
}

func (c *Client) version(component types.VersionInfo) (string, error) {
	cmd := commands.QueryCommandVersion{
		Component: component,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return "", err
	}
	if len(resp) == 0 {
		return "", nil
	}
	return strings.TrimSpace(resp[0]), nil
}

type QueryInfoCommand struct {
	QueryBuilder
}

func (b *QueryBuilder) Info() *QueryInfoCommand {
	return &QueryInfoCommand{
		QueryBuilder: *b,
	}
}

func (b *QueryInfoCommand) sendInfo(component types.InfoComponent) ([]string, error) {
	cmd := commands.QueryCommandInfo{
		Component: component,
	}
	return b.client.Send(cmd)
}

// Generic retrieves a list of the available channels.
func (b *QueryInfoCommand) Generic() ([]responses.QueryChannelInfo, error) {
	resp, err := b.sendInfo("")
	if err != nil {
		return nil, err
	}
	return responses.ResponseToQueryChannelInfo(resp)
}

// Template gets information about the specified template.
//
// WARNING: This command does not return what it states as of CasparCG 2.5.0
//
// https://github.com/CasparCG/server/issues/1151
func (b *QueryInfoCommand) Template(template string) (responses.QueryChannelInfo, error) {
	cmd := commands.QueryCommandInfoTemplate{
		Template: template,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return responses.QueryChannelInfo{}, err
	}
	return responses.PartsToQueryChannelInfo(resp)
}

// Config gets the contents of the configuration used.
func (b *QueryInfoCommand) Config() (responses.CasparConfig, error) {
	resp, err := b.sendInfo(types.InfoComponentConfig)
	if err != nil {
		return responses.CasparConfig{}, err
	}
	var config responses.CasparConfig
	err = xml.Unmarshal([]byte(strings.Join(resp, "\n")), &config)
	if err != nil {
		return responses.CasparConfig{}, err
	}
	return config, nil
}

// Paths gets information about the paths used.
func (b *QueryInfoCommand) Paths() (responses.Paths, error) {
	resp, err := b.sendInfo(types.InfoComponentPaths)
	if err != nil {
		return responses.Paths{}, err
	}
	var paths responses.Paths
	err = xml.Unmarshal([]byte(strings.Join(resp, "\n")), &paths)
	if err != nil {
		return responses.Paths{}, err
	}
	return paths, nil
}

// System gets system information like OS, CPU and library version numbers.
//
// WARNING: This command does not return what it states as of CasparCG 2.5.0
//
// https://github.com/CasparCG/server/issues/1151
func (b *QueryInfoCommand) System() ([]responses.QueryChannelInfo, error) {
	resp, err := b.sendInfo(types.InfoComponentSystem)
	if err != nil {
		return nil, err
	}
	return responses.ResponseToQueryChannelInfo(resp)

}

// Server gets detailed information about all channels.
func (b *QueryInfoCommand) Server() ([]responses.QueryChannelInfo, error) {
	resp, err := b.sendInfo(types.InfoComponentServer)
	if err != nil {
		return nil, err
	}
	return responses.ResponseToQueryChannelInfo(resp)
}

// Queues gets detailed information about all AMCP Command Queues.
//
// WARNING: This command does not return what it states as of CasparCG 2.5.0
//
// https://github.com/CasparCG/server/issues/1151
func (b *QueryInfoCommand) Queues() ([]responses.QueryChannelInfo, error) {
	resp, err := b.sendInfo(types.InfoComponentQueues)
	if err != nil {
		return nil, err
	}
	return responses.ResponseToQueryChannelInfo(resp)
}

// Threads lists all known threads in the server.
//
// WARNING: This command does not return what it states as of CasparCG 2.5.0
//
// https://github.com/CasparCG/server/issues/1151
func (b *QueryInfoCommand) Threads() ([]responses.QueryChannelInfo, error) {
	resp, err := b.sendInfo(types.InfoComponentThreads)
	if err != nil {
		return nil, err
	}
	return responses.ResponseToQueryChannelInfo(resp)
}
