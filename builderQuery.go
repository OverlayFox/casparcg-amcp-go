package casparcg

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
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

// Generic retrieves a list of the available channels.
func (b *QueryInfoCommand) Generic() ([]responses.QueryChannelInfo, error) {
	resp, err := b.sendInfo("")
	if err != nil {
		return nil, err
	}
	channelInfo := make([]responses.QueryChannelInfo, 0, len(resp))
	for _, line := range resp {
		parts := strings.Split(line, " ")
		info, err := responses.PartsToQueryChannelInfo(parts)
		if err != nil {
			return nil, err
		}
		channelInfo = append(channelInfo, info)
	}
	return channelInfo, nil
}

// Template gets information about the specified template.
func (b *QueryInfoCommand) Template(template string) (responses.QueryChannelInfo, error) {
	cmd := commands.QueryCommandInfoTemplate{
		Template: template,
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.QueryChannelInfo{}, err
	}

	parts := strings.Split(data[0], " ")
	if len(parts) != 3 {
		return responses.QueryChannelInfo{}, fmt.Errorf("unexpected format for template info: %s", data[0])
	}

	videoChannel, err := strconv.Atoi(parts[0])
	if err != nil {
		return responses.QueryChannelInfo{}, fmt.Errorf("invalid video channel in template info: %s", parts[0])
	}

	return responses.QueryChannelInfo{
		ChannelIndex: videoChannel,
		VideoMode:    types.VideoMode(parts[1]),
		Status:       parts[2],
	}, nil
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

// // System gets system information like OS, CPU and library version numbers.
// func (b *QueryInfoCommand) System() (responses.QueryChannelInfo, error) {
// 	data, err := b.info(types.InfoComponentSystem)
// 	if err != nil {
// 		systemInfo, ok := data.(responses.QueryChannelInfo)
// 		if !ok {
// 			return responses.QueryChannelInfo{}, fmt.Errorf("unexpected data type for system info: %T", data)
// 		}
// 		return systemInfo, nil
// 	}
// 	return responses.QueryChannelInfo{}, err
// }

// func (b *QueryInfoCommand) Server() (responses.QueryChannelInfo, error) {
// 	data, err := b.info(types.InfoComponentServer)
// 	if err != nil {
// 		return responses.QueryChannelInfo{}, err
// 	}
// 	serverInfo, ok := data.(responses.QueryChannelInfo)
// 	if !ok {
// 		return responses.QueryChannelInfo{}, fmt.Errorf("unexpected data type for server info: %T", data)
// 	}
// 	return serverInfo, nil
// }

// func (b *QueryInfoCommand) Queues() (responses.QueryChannelInfo, error) {
// 	data, err := b.info(types.InfoComponentQueues)
// 	if err != nil {
// 		return responses.QueryChannelInfo{}, err
// 	}
// 	queuesInfo, ok := data.(responses.QueryChannelInfo)
// 	if !ok {
// 		return responses.QueryChannelInfo{}, fmt.Errorf("unexpected data type for queues info: %T", data)
// 	}
// 	return queuesInfo, nil
// }

// func (b *QueryInfoCommand) Threads() (responses.QueryChannelInfo, error) {
// 	data, err := b.info(types.InfoComponentThreads)
// 	if err != nil {
// 		return responses.QueryChannelInfo{}, err
// 	}
// 	threadsInfo, ok := data.(responses.QueryChannelInfo)
// 	if !ok {
// 		return responses.QueryChannelInfo{}, fmt.Errorf("unexpected data type for threads info: %T", data)
// 	}
// 	return threadsInfo, nil
// }

// Channel gets information about a channel or layer.
func (b *QueryInfoCommand) Channel(videoChannel int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        nil,
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	var infoChannel responses.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	return infoChannel, nil
}

func (b *QueryInfoCommand) ChannelLayer(videoChannel int, layer int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        &layer,
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	var infoChannel responses.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	return infoChannel, nil
}

// InfoChannelDelay gets delay information.
func (b *QueryInfoCommand) InfoChannelDelay(videoChannel int, layer *int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        layer,
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	var infoChannel responses.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	return infoChannel, nil
}

func (b *QueryInfoCommand) ChannelLayerDelay(videoChannel int, layer int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        &layer,
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	var infoChannel responses.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return responses.InfoChannel{}, err
	}

	return infoChannel, nil
}

func (b *QueryInfoCommand) sendInfo(component types.InfoComponent) ([]string, error) {
	cmd := commands.QueryCommandInfo{
		Component: component,
	}
	return b.client.Send(cmd)

	// switch component {
	// case types.InfoComponentConfig:

	// case types.InfoComponentSystem, types.InfoComponentServer, types.InfoComponentQueues, types.InfoComponentThreads:
	// 	parts := strings.Split(fullXML, " ")
	// 	return responses.PartsToQueryChannelInfo(parts)

	// default:
	// 	return nil, fmt.Errorf("unknown info component: %s", component)
	// }
}
