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

// sendCommand abstracts sending a command that does not expect a response value.
func (b *Client) sendCommand(cmd interface{ String() string }) error {
	_, err := b.Send(cmd)
	return err
}

// LogLevel sets the log level.
func (c *Client) LogLevel(level types.LogLevel) error {
	cmd := commands.DirectCommandLogLevel{
		Level: level,
	}
	return c.sendCommand(cmd)
}

// Ping sends a ping command.
func (c *Client) Ping(token *string) (string, error) {
	cmd := commands.DirectCommandPing{
		Token: token,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return "", err
	}
	return data[0], nil
}

// Bye closes the connection.
func (c *Client) Bye() error {
	cmd := commands.CommandBye{}
	_, err := c.Send(cmd)
	return err
}

// Kill kills the server.
func (c *Client) Kill() error {
	cmd := commands.CommandKill{}
	_, err := c.Send(cmd)
	return err
}

// Restart restarts the server.
func (c *Client) Restart() error {
	cmd := commands.CommandRestart{}
	_, err := c.Send(cmd)
	return err
}

var reCINF = regexp.MustCompile(`^"?([^"]+)"?\s+(\S+)\s+(\d+)\s+(\d+)\s+(\d+)\s+([\d/]+)$`)

// CINF returns information about a media file.
func (c *Client) CINF(filename string) (responses.CINF, error) {
	cmd := commands.QueryCommandCINF{
		Filename: filename,
	}
	resp, err := c.Send(cmd)
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
func (c *Client) CLS(directory *string) ([]responses.CINF, error) {
	cmd := commands.QueryCommandCLS{
		Directory: directory,
	}
	resp, err := c.Send(cmd)
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
// TODO: implement DTO for FLS response
func (c *Client) FLS() ([]string, error) {
	cmd := commands.QueryCommandFLS{}
	return c.Send(cmd)
}

// TLS lists template files.
func (c *Client) TLS(directory string) ([]string, error) {
	cmd := commands.QueryCommandTLS{
		Directory: directory,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type VersionBuilder struct {
	client *Client
}

// Version returns the version of specified component.
func (c *Client) Version() *VersionBuilder {
	return &VersionBuilder{
		client: c,
	}
}

func (b *VersionBuilder) Generic() (string, error) {
	return b.client.version("")
}

func (b *VersionBuilder) Server() (string, error) {
	return b.client.version(types.VersionInfoServer)
}

func (b *VersionBuilder) Flash() (string, error) {
	return b.client.version(types.VersionInfoFlash)
}

func (b *VersionBuilder) TemplateHost() (string, error) {
	return b.client.version(types.VersionInfoTemplateHost)
}

func (b *VersionBuilder) CEF() (string, error) {
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

type InfoBuilder struct {
	client *Client
}

func (c *Client) Info() *InfoBuilder {
	return &InfoBuilder{
		client: c,
	}
}

func (b *InfoBuilder) Generic() (types.InfoComponent, error) {
	data, err := b.info("")
	if err != nil {
		config, ok := data.(types.InfoComponent)
		if !ok {
			return "", fmt.Errorf("unexpected data type for config info: %T", data)
		}
		return config, nil
	}
	return "", err
}

func (b *InfoBuilder) Config() (responses.CasparConfig, error) {
	data, err := b.info(types.InfoComponentConfig)
	if err != nil {
		config, ok := data.(responses.CasparConfig)
		if !ok {
			return responses.CasparConfig{}, fmt.Errorf("unexpected data type for config info: %T", data)
		}
		return config, nil
	}
	return responses.CasparConfig{}, err
}

func (b *InfoBuilder) Paths() (responses.Paths, error) {
	data, err := b.info(types.InfoComponentPaths)
	if err != nil {
		paths, ok := data.(responses.Paths)
		if !ok {
			return responses.Paths{}, fmt.Errorf("unexpected data type for paths info: %T", data)
		}
		return paths, nil
	}
	return responses.Paths{}, err
}

func (b *InfoBuilder) System() (responses.GenericInfo, error) {
	data, err := b.info(types.InfoComponentSystem)
	if err != nil {
		systemInfo, ok := data.(responses.GenericInfo)
		if !ok {
			return responses.GenericInfo{}, fmt.Errorf("unexpected data type for system info: %T", data)
		}
		return systemInfo, nil
	}
	return responses.GenericInfo{}, err
}

func (b *InfoBuilder) Server() (responses.GenericInfo, error) {
	data, err := b.info(types.InfoComponentServer)
	if err != nil {
		return responses.GenericInfo{}, err
	}
	serverInfo, ok := data.(responses.GenericInfo)
	if !ok {
		return responses.GenericInfo{}, fmt.Errorf("unexpected data type for server info: %T", data)
	}
	return serverInfo, nil
}

func (b *InfoBuilder) Queues() (responses.GenericInfo, error) {
	data, err := b.info(types.InfoComponentQueues)
	if err != nil {
		return responses.GenericInfo{}, err
	}
	queuesInfo, ok := data.(responses.GenericInfo)
	if !ok {
		return responses.GenericInfo{}, fmt.Errorf("unexpected data type for queues info: %T", data)
	}
	return queuesInfo, nil
}

func (b *InfoBuilder) Threads() (responses.GenericInfo, error) {
	data, err := b.info(types.InfoComponentThreads)
	if err != nil {
		return responses.GenericInfo{}, err
	}
	threadsInfo, ok := data.(responses.GenericInfo)
	if !ok {
		return responses.GenericInfo{}, fmt.Errorf("unexpected data type for threads info: %T", data)
	}
	return threadsInfo, nil
}

func (b *InfoBuilder) info(component types.InfoComponent) (any, error) {
	cmd := commands.QueryCommandInfo{
		Component: component,
	}
	resp, err := b.client.Send(cmd)
	if err != nil {
		return nil, err
	}

	fullXML := strings.Join(resp, "\n")
	switch component {
	case types.InfoComponentConfig:
		var config responses.CasparConfig
		err := xml.Unmarshal([]byte(fullXML), &config)
		if err != nil {
			return nil, err
		}
		return config, nil

	case types.InfoComponentPaths:
		var paths responses.Paths
		err := xml.Unmarshal([]byte(fullXML), &paths)
		if err != nil {
			return nil, err
		}
		return paths, nil

	case types.InfoComponentSystem, types.InfoComponentServer, types.InfoComponentQueues, types.InfoComponentThreads:
		parts := strings.Split(fullXML, " ")
		if len(parts) != 3 {
			return nil, fmt.Errorf("unexpected format for '%s' info: %s", component, fullXML)
		}

		videoChannel, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid video channel in '%s' info: %s", component, parts[0])
		}

		systemInfo := responses.GenericInfo{
			VideoChannel: videoChannel,
			VideoMode:    types.VideoMode(parts[1]),
			Status:       parts[2],
		}
		return systemInfo, nil

	default:
		return nil, fmt.Errorf("unknown info component: %s", component)
	}
}

// Channel gets information about a channel or layer.
func (b *InfoBuilder) Channel(videoChannel int) (responses.InfoChannel, error) {
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

func (b *InfoBuilder) ChannelLayer(videoChannel int, layer int) (responses.InfoChannel, error) {
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

// Template gets information about the specified template.
func (b *InfoBuilder) Template(template string) (responses.GenericInfo, error) {
	cmd := commands.QueryCommandInfoTemplate{
		Template: template,
	}
	data, err := b.client.Send(cmd)
	if err != nil {
		return responses.GenericInfo{}, err
	}

	parts := strings.Split(data[0], " ")
	if len(parts) != 3 {
		return responses.GenericInfo{}, fmt.Errorf("unexpected format for template info: %s", data[0])
	}

	videoChannel, err := strconv.Atoi(parts[0])
	if err != nil {
		return responses.GenericInfo{}, fmt.Errorf("invalid video channel in template info: %s", parts[0])
	}

	return responses.GenericInfo{
		VideoChannel: videoChannel,
		VideoMode:    types.VideoMode(parts[1]),
		Status:       parts[2],
	}, nil
}

// InfoChannelDelay gets delay information.
func (b *InfoBuilder) InfoChannelDelay(videoChannel int, layer *int) (responses.InfoChannel, error) {
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

func (b *InfoBuilder) ChannelLayerDelay(videoChannel int, layer int) (responses.InfoChannel, error) {
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
