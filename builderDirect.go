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

// Direct command methods on Client for commands that don't require a builder

// SWAP swaps layers between channels.
func (c *Client) SWAP(channel1, channel2 int, layer1, layer2 *int, transform bool) error {
	cmd := commands.CommandSwap{
		VideoChannel1: channel1,
		Layer1:        layer1,
		VideoChannel2: channel2,
		Layer2:        layer2,
		Transform:     transform,
	}
	_, err := c.Send(cmd)
	return err
}

// ADD adds a consumer to the specified video channel.
func (c *Client) ADD(
	videoChannel int,
	consumerIdx *int,
	consumerName string,
	parameters map[string]string,
) error {
	cmd := commands.CommandAdd{
		VideoChannel: videoChannel,
		ConsumerIdx:  consumerIdx,
		ConsumerName: consumerName,
		Parameters:   parameters,
	}
	_, err := c.Send(cmd)
	return err
}

// REMOVE removes a consumer from the specified video channel.
func (c *Client) REMOVE(videoChannel int, consumerIdx *int, parameters *map[string]string) error {
	cmd := commands.CommandRemove{
		VideoChannel: videoChannel,
		ConsumerIdx:  consumerIdx,
		Parameters:   parameters,
	}
	_, err := c.Send(cmd)
	return err
}

// PRINT sends a print command for the specified video channel.
func (c *Client) PRINT(videoChannel int) error {
	cmd := commands.CommandPrint{
		VideoChannel: videoChannel,
	}
	_, err := c.Send(cmd)
	return err
}

// LOGLEVEL sets the log level.
func (c *Client) LOGLEVEL(level types.LogLevel) error {
	cmd := commands.CommandLogLevel{
		Level: level,
	}
	_, err := c.Send(cmd)
	return err
}

// SET changes the value of a channel variable.
func (c *Client) SET(videoChannel int, variable types.SetVariable, value string) error {
	cmd := commands.CommandSet{
		VideoChannel: videoChannel,
		Variable:     variable,
		Value:        value,
	}
	_, err := c.Send(cmd)
	return err
}

// LOCK performs a lock operation on the specified channel.
func (c *Client) LOCK(videoChannel int, action types.LockAction, secret *string) error {
	cmd := commands.CommandLock{
		VideoChannel: videoChannel,
		Action:       action,
		Secret:       secret,
	}
	_, err := c.Send(cmd)
	return err
}

// PING sends a ping command.
func (c *Client) PING(token string) (string, error) {
	cmd := commands.CommandPing{
		Token: token,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return "", err
	}
	return data[0], nil
}

// BYE closes the connection.
func (c *Client) BYE() error {
	cmd := commands.CommandBye{}
	_, err := c.Send(cmd)
	return err
}

// KILL kills the server.
func (c *Client) KILL() error {
	cmd := commands.CommandKill{}
	_, err := c.Send(cmd)
	return err
}

// RESTART restarts the server.
func (c *Client) RESTART() error {
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

// VERSION returns the version of specified component.
func (c *Client) VERSION() (string, error) {
	return c.version("")
}

func (c *Client) VERSIONSERVER() (string, error) {
	return c.version(types.VersionInfoServer)
}

func (c *Client) VERSIONFLASH() (string, error) {
	return c.version(types.VersionInfoFlash)
}

func (c *Client) VERSIONTEMPLATEHOST() (string, error) {
	return c.version(types.VersionInfoTemplateHost)
}

func (c *Client) VERSIONCEF() (string, error) {
	return c.version(types.VersionInfoCEF)
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

func (c *Client) INFO() (types.InfoComponent, error) {
	data, err := c.info("")
	if err != nil {
		config, ok := data.(types.InfoComponent)
		if !ok {
			return "", fmt.Errorf("unexpected data type for config info: %T", data)
		}
		return config, nil
	}
	return "", err
}

func (c *Client) INFOCONFIG() (responses.CasparConfig, error) {
	data, err := c.info(types.InfoComponentConfig)
	if err != nil {
		config, ok := data.(responses.CasparConfig)
		if !ok {
			return responses.CasparConfig{}, fmt.Errorf("unexpected data type for config info: %T", data)
		}
		return config, nil
	}
	return responses.CasparConfig{}, err
}

func (c *Client) INFOPATHS() (responses.Paths, error) {
	data, err := c.info(types.InfoComponentPaths)
	if err != nil {
		paths, ok := data.(responses.Paths)
		if !ok {
			return responses.Paths{}, fmt.Errorf("unexpected data type for paths info: %T", data)
		}
		return paths, nil
	}
	return responses.Paths{}, err
}

func (c *Client) INFOSYSTEM() (responses.GenericInfo, error) {
	data, err := c.info(types.InfoComponentSystem)
	if err != nil {
		systemInfo, ok := data.(responses.GenericInfo)
		if !ok {
			return responses.GenericInfo{}, fmt.Errorf("unexpected data type for system info: %T", data)
		}
		return systemInfo, nil
	}
	return responses.GenericInfo{}, err
}

func (c *Client) INFOSERVER() (responses.GenericInfo, error) {
	data, err := c.info(types.InfoComponentServer)
	if err != nil {
		systemInfo, ok := data.(responses.GenericInfo)
		if !ok {
			return responses.GenericInfo{}, fmt.Errorf("unexpected data type for server info: %T", data)
		}
		return systemInfo, nil
	}
	return responses.GenericInfo{}, err
}

func (c *Client) INFOQUEUES() (responses.GenericInfo, error) {
	data, err := c.info(types.InfoComponentQueues)
	if err != nil {
		systemInfo, ok := data.(responses.GenericInfo)
		if !ok {
			return responses.GenericInfo{}, fmt.Errorf("unexpected data type for queues info: %T", data)
		}
		return systemInfo, nil
	}
	return responses.GenericInfo{}, err
}

func (c *Client) INFOTHREADS() (responses.GenericInfo, error) {
	data, err := c.info(types.InfoComponentThreads)
	if err != nil {
		systemInfo, ok := data.(responses.GenericInfo)
		if !ok {
			return responses.GenericInfo{}, fmt.Errorf("unexpected data type for threads info: %T", data)
		}
		return systemInfo, nil
	}
	return responses.GenericInfo{}, err
}

func (c *Client) info(component types.InfoComponent) (any, error) {
	cmd := commands.QueryCommandInfo{
		Component: component,
	}
	resp, err := c.Send(cmd)
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

// INFOCHANNEL gets information about a channel or layer.
func (c *Client) INFOCHANNEL(videoChannel int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        nil,
	}
	data, err := c.Send(cmd)
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

func (c *Client) INFOCHANNELLAYER(videoChannel int, layer int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        &layer,
	}
	data, err := c.Send(cmd)
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

// INFOTEMPLATE gets information about the specified template.
func (c *Client) INFOTEMPLATE(template string) (responses.GenericInfo, error) {
	cmd := commands.QueryCommandInfoTemplate{
		Template: template,
	}
	data, err := c.Send(cmd)
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

// INFOCHANNELDELAY gets delay information.
func (c *Client) INFOCHANNELDELAY(videoChannel int, layer *int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        layer,
	}
	data, err := c.Send(cmd)
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

func (c *Client) INFOCHANNELLAYERDELAY(videoChannel int, layer int) (responses.InfoChannel, error) {
	cmd := commands.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        &layer,
	}
	data, err := c.Send(cmd)
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
