package casparcg

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/returns"
)

// CGBuilder provides a fluent interface for building CG (template) commands.
type CGBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

// CG creates a new CG command builder for the specified channel and layer
// Example: client.CG(1, 12).STOP(2).
func (c *Client) CG(videoChannel, layer int) *CGBuilder {
	return &CGBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// ADD prepares a template for displaying.
func (b *CGBuilder) ADD(cgLayer int, template string, playOnLoad bool, data *string) error {
	cmd := types.TemplateCommandCGAdd{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer:    cgLayer,
		Template:   template,
		PlayOnLoad: playOnLoad,
		Data:       data,
	}
	_, err := b.client.Send(cmd)
	return err
}

// PLAY plays and displays the template in the specified layer.
func (b *CGBuilder) PLAY(cgLayer int) error {
	cmd := types.TemplateCommandCGPlay{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// STOP stops the template in the specified layer.
func (b *CGBuilder) STOP(cgLayer int) error {
	cmd := types.TemplateCommandCGStop{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// NEXT triggers a "continue" in the template.
func (b *CGBuilder) NEXT(cgLayer int) error {
	cmd := types.TemplateCommandCGNext{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// REMOVE removes the template from the specified layer.
func (b *CGBuilder) REMOVE(cgLayer int) error {
	cmd := types.TemplateCommandCGRemove{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// CLEAR removes all templates on the video layer.
func (b *CGBuilder) CLEAR() error {
	cmd := types.TemplateCommandCGClear{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// UPDATE sends new data to the template on specified layer.
func (b *CGBuilder) UPDATE(cgLayer int, data string) error {
	cmd := types.TemplateCommandCGUpdate{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
		Data:    data,
	}
	_, err := b.client.Send(cmd)
	return err
}

// INVOKE invokes the given method on the template.
func (b *CGBuilder) INVOKE(cgLayer int, method string) error {
	cmd := types.TemplateCommandCGInvoke{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
		Method:  method,
	}
	_, err := b.client.Send(cmd)
	return err
}

// INFO retrieves information about the template on the specified layer.
func (b *CGBuilder) INFO(cgLayer *int) error {
	cmd := types.TemplateCommandCGInfo{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	_, err := b.client.Send(cmd)
	return err
}

// LayerBuilder provides a fluent interface for building layer-based commands.
type LayerBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

// Layer creates a new layer command builder for the specified channel and layer
// Example: client.Layer(1, 10).PLAY("myclip", nil).
func (c *Client) Layer(videoChannel, layer int) *LayerBuilder {
	return &LayerBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// LOAD loads a clip to the layer.
func (b *LayerBuilder) LOAD(clip string, parameters *map[string]string) error {
	cmd := types.CommandLoad{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Clip:       clip,
		Parameters: parameters,
	}
	_, err := b.client.Send(cmd)
	return err
}

// PLAY plays content on the layer.
func (b *LayerBuilder) PLAY(clip *string, parameters *map[string]string) error {
	cmd := types.CommandPlay{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Clip:       clip,
		Parameters: parameters,
	}
	_, err := b.client.Send(cmd)
	return err
}

// PAUSE pauses playback on the layer.
func (b *LayerBuilder) PAUSE() error {
	cmd := types.CommandPause{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// RESUME resumes playback on the layer.
func (b *LayerBuilder) RESUME() error {
	cmd := types.CommandResume{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// STOP stops playback on the layer.
func (b *LayerBuilder) STOP() error {
	cmd := types.CommandStop{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// CLEAR clears the layer.
func (b *LayerBuilder) CLEAR() error {
	cmd := types.CommandClear{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	_, err := b.client.Send(cmd)
	return err
}

// CALL calls a function on the layer.
func (b *LayerBuilder) CALL(params map[string]string) error {
	cmd := types.CommandCall{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Params: params,
	}
	_, err := b.client.Send(cmd)
	return err
}

// Direct command methods on Client for commands that don't require a builder

// SWAP swaps layers between channels.
func (c *Client) SWAP(channel1, channel2 int, layer1, layer2 *int, transform bool) error {
	cmd := types.CommandSwap{
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
	cmd := types.CommandAdd{
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
	cmd := types.CommandRemove{
		VideoChannel: videoChannel,
		ConsumerIdx:  consumerIdx,
		Parameters:   parameters,
	}
	_, err := c.Send(cmd)
	return err
}

// PRINT sends a print command for the specified video channel.
func (c *Client) PRINT(videoChannel int) error {
	cmd := types.CommandPrint{
		VideoChannel: videoChannel,
	}
	_, err := c.Send(cmd)
	return err
}

// LOGLEVEL sets the log level.
func (c *Client) LOGLEVEL(level types.LogLevel) error {
	cmd := types.CommandLogLevel{
		Level: level,
	}
	_, err := c.Send(cmd)
	return err
}

// SET changes the value of a channel variable.
func (c *Client) SET(videoChannel int, variable types.SetVariable, value string) error {
	cmd := types.CommandSet{
		VideoChannel: videoChannel,
		Variable:     variable,
		Value:        value,
	}
	_, err := c.Send(cmd)
	return err
}

// LOCK performs a lock operation on the specified channel.
func (c *Client) LOCK(videoChannel int, action types.LockAction, secret *string) error {
	cmd := types.CommandLock{
		VideoChannel: videoChannel,
		Action:       action,
		Secret:       secret,
	}
	_, err := c.Send(cmd)
	return err
}

// PING sends a ping command.
func (c *Client) PING(token string) (string, error) {
	cmd := types.CommandPing{
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
	cmd := types.CommandBye{}
	_, err := c.Send(cmd)
	return err
}

// KILL kills the server.
func (c *Client) KILL() error {
	cmd := types.CommandKill{}
	_, err := c.Send(cmd)
	return err
}

// RESTART restarts the server.
func (c *Client) RESTART() error {
	cmd := types.CommandRestart{}
	_, err := c.Send(cmd)
	return err
}

// Query command methods

var reCINF = regexp.MustCompile(`^"?([^"]+)"?\s+(\S+)\s+(\d+)\s+(\d+)\s+(\d+)\s+([\d/]+)$`)

// CINF returns information about a media file.
func (c *Client) CINF(filename string) (returns.CINF, error) {
	cmd := types.QueryCommandCINF{
		Filename: filename,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return returns.CINF{}, err
	}

	matches := reCINF.FindStringSubmatch(resp[0])
	cinf, err := matchesToCINF(matches)
	if err != nil {
		return returns.CINF{}, err
	}

	return cinf, nil
}

// CLS lists media files in the media folder.
func (c *Client) CLS(directory *string) ([]returns.CINF, error) {
	cmd := types.QueryCommandCLS{
		Directory: directory,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return nil, err
	}

	cls := []returns.CINF{}
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

func matchesToCINF(matches []string) (returns.CINF, error) {
	if matches == nil || len(matches) != 7 {
		return returns.CINF{}, fmt.Errorf("unexpected format for CINF response: %s", matches)
	}

	cinfSize, err := strconv.Atoi(strings.TrimSpace(matches[3]))
	if err != nil {
		return returns.CINF{}, fmt.Errorf("invalid file size in CINF response: %s", matches[3])
	}

	cinfLastModified, err := time.Parse("20060102150405", strings.TrimSpace(matches[4]))
	if err != nil {
		return returns.CINF{}, fmt.Errorf("invalid last modified date in CINF response: %s", matches[4])
	}

	cinfFrameCount, err := strconv.Atoi(strings.TrimSpace(matches[5]))
	if err != nil {
		return returns.CINF{}, fmt.Errorf("invalid frame count in CINF response: %s", matches[5])
	}

	cinfFrameRate, err := types.StringToFrameRate(strings.TrimSpace(matches[6]))
	if err != nil {
		return returns.CINF{}, fmt.Errorf("invalid frame rate in CINF response: %s", matches[6])
	}

	return returns.CINF{
		Filename:     strings.TrimSpace(matches[1]),
		Type:         types.MediaTypes(strings.TrimSpace(matches[2])),
		FileSize:     int64(cinfSize),
		LastModified: cinfLastModified,
		FrameCount:   cinfFrameCount,
		FrameRate:    cinfFrameRate,
	}, nil
}

// FLS lists all fonts
// TODO: implement DTO for FLS response
func (c *Client) FLS() ([]string, error) {
	cmd := types.QueryCommandFLS{}
	return c.Send(cmd)
}

// TLS lists template files.
func (c *Client) TLS(directory string) ([]string, error) {
	cmd := types.QueryCommandTLS{
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
	cmd := types.QueryCommandVersion{
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

func (c *Client) INFOCONFIG() (returns.CasparConfig, error) {
	data, err := c.info(types.InfoComponentConfig)
	if err != nil {
		config, ok := data.(returns.CasparConfig)
		if !ok {
			return returns.CasparConfig{}, fmt.Errorf("unexpected data type for config info: %T", data)
		}
		return config, nil
	}
	return returns.CasparConfig{}, err
}

func (c *Client) INFOPATHS() (returns.Paths, error) {
	data, err := c.info(types.InfoComponentPaths)
	if err != nil {
		paths, ok := data.(returns.Paths)
		if !ok {
			return returns.Paths{}, fmt.Errorf("unexpected data type for paths info: %T", data)
		}
		return paths, nil
	}
	return returns.Paths{}, err
}

func (c *Client) INFOSYSTEM() (returns.GenericInfo, error) {
	data, err := c.info(types.InfoComponentSystem)
	if err != nil {
		systemInfo, ok := data.(returns.GenericInfo)
		if !ok {
			return returns.GenericInfo{}, fmt.Errorf("unexpected data type for system info: %T", data)
		}
		return systemInfo, nil
	}
	return returns.GenericInfo{}, err
}

func (c *Client) INFOSERVER() (returns.GenericInfo, error) {
	data, err := c.info(types.InfoComponentServer)
	if err != nil {
		systemInfo, ok := data.(returns.GenericInfo)
		if !ok {
			return returns.GenericInfo{}, fmt.Errorf("unexpected data type for server info: %T", data)
		}
		return systemInfo, nil
	}
	return returns.GenericInfo{}, err
}

func (c *Client) INFOQUEUES() (returns.GenericInfo, error) {
	data, err := c.info(types.InfoComponentQueues)
	if err != nil {
		systemInfo, ok := data.(returns.GenericInfo)
		if !ok {
			return returns.GenericInfo{}, fmt.Errorf("unexpected data type for queues info: %T", data)
		}
		return systemInfo, nil
	}
	return returns.GenericInfo{}, err
}

func (c *Client) INFOTHREADS() (returns.GenericInfo, error) {
	data, err := c.info(types.InfoComponentThreads)
	if err != nil {
		systemInfo, ok := data.(returns.GenericInfo)
		if !ok {
			return returns.GenericInfo{}, fmt.Errorf("unexpected data type for threads info: %T", data)
		}
		return systemInfo, nil
	}
	return returns.GenericInfo{}, err
}

func (c *Client) info(component types.InfoComponent) (any, error) {
	cmd := types.QueryCommandInfo{
		Component: component,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return nil, err
	}

	fullXML := strings.Join(resp, "\n")
	switch component {
	case types.InfoComponentConfig:
		var config returns.CasparConfig
		err := xml.Unmarshal([]byte(fullXML), &config)
		if err != nil {
			return nil, err
		}
		return config, nil

	case types.InfoComponentPaths:
		var paths returns.Paths
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

		systemInfo := returns.GenericInfo{
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
func (c *Client) INFOCHANNEL(videoChannel int) (returns.InfoChannel, error) {
	cmd := types.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        nil,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	var infoChannel returns.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	return infoChannel, nil
}

func (c *Client) INFOCHANNELLAYER(videoChannel int, layer int) (returns.InfoChannel, error) {
	cmd := types.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        &layer,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	var infoChannel returns.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	return infoChannel, nil
}

// INFOTEMPLATE gets information about the specified template.
func (c *Client) INFOTEMPLATE(template string) (returns.GenericInfo, error) {
	cmd := types.QueryCommandInfoTemplate{
		Template: template,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return returns.GenericInfo{}, err
	}

	parts := strings.Split(data[0], " ")
	if len(parts) != 3 {
		return returns.GenericInfo{}, fmt.Errorf("unexpected format for template info: %s", data[0])
	}

	videoChannel, err := strconv.Atoi(parts[0])
	if err != nil {
		return returns.GenericInfo{}, fmt.Errorf("invalid video channel in template info: %s", parts[0])
	}

	return returns.GenericInfo{
		VideoChannel: videoChannel,
		VideoMode:    types.VideoMode(parts[1]),
		Status:       parts[2],
	}, nil
}

// INFOCHANNELDELAY gets delay information.
func (c *Client) INFOCHANNELDELAY(videoChannel int, layer *int) (returns.InfoChannel, error) {
	cmd := types.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        layer,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	var infoChannel returns.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	return infoChannel, nil
}

func (c *Client) INFOCHANNELLAYERDELAY(videoChannel int, layer int) (returns.InfoChannel, error) {
	cmd := types.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        &layer,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	var infoChannel returns.InfoChannel
	err = xml.Unmarshal([]byte(strings.Join(data, "\n")), &infoChannel)
	if err != nil {
		return returns.InfoChannel{}, err
	}

	return infoChannel, nil
}
