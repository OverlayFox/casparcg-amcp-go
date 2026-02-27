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

// CGBuilder provides a fluent interface for building CG (template) commands
type CGBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

// CG creates a new CG command builder for the specified channel and layer
// Example: client.CG(1, 12).STOP(2)
func (c *Client) CG(videoChannel, layer int) *CGBuilder {
	return &CGBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// ADD prepares a template for displaying
func (b *CGBuilder) ADD(cgLayer int, template string, playOnLoad bool, data *string) (*Response, error) {
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
	return b.client.Send(cmd)
}

// PLAY plays and displays the template in the specified layer
func (b *CGBuilder) PLAY(cgLayer int) (*Response, error) {
	cmd := types.TemplateCommandCGPlay{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	return b.client.Send(cmd)
}

// STOP stops the template in the specified layer
func (b *CGBuilder) STOP(cgLayer int) (*Response, error) {
	cmd := types.TemplateCommandCGStop{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	return b.client.Send(cmd)
}

// NEXT triggers a "continue" in the template
func (b *CGBuilder) NEXT(cgLayer int) (*Response, error) {
	cmd := types.TemplateCommandCGNext{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	return b.client.Send(cmd)
}

// REMOVE removes the template from the specified layer
func (b *CGBuilder) REMOVE(cgLayer int) (*Response, error) {
	cmd := types.TemplateCommandCGRemove{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	return b.client.Send(cmd)
}

// CLEAR removes all templates on the video layer
func (b *CGBuilder) CLEAR() (*Response, error) {
	cmd := types.TemplateCommandCGClear{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	return b.client.Send(cmd)
}

// UPDATE sends new data to the template on specified layer
func (b *CGBuilder) UPDATE(cgLayer int, data string) (*Response, error) {
	cmd := types.TemplateCommandCGUpdate{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
		Data:    data,
	}
	return b.client.Send(cmd)
}

// INVOKE invokes the given method on the template
func (b *CGBuilder) INVOKE(cgLayer int, method string) (*Response, error) {
	cmd := types.TemplateCommandCGInvoke{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
		Method:  method,
	}
	return b.client.Send(cmd)
}

// INFO retrieves information about the template on the specified layer
func (b *CGBuilder) INFO(cgLayer *int) (*Response, error) {
	cmd := types.TemplateCommandCGInfo{
		TemplateCommandCG: types.TemplateCommandCG{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		CgLayer: cgLayer,
	}
	return b.client.Send(cmd)
}

// LayerBuilder provides a fluent interface for building layer-based commands
type LayerBuilder struct {
	client       *Client
	videoChannel int
	layer        int
}

// Layer creates a new layer command builder for the specified channel and layer
// Example: client.Layer(1, 10).PLAY("myclip", nil)
func (c *Client) Layer(videoChannel, layer int) *LayerBuilder {
	return &LayerBuilder{
		client:       c,
		videoChannel: videoChannel,
		layer:        layer,
	}
}

// LOAD loads a clip to the layer
func (b *LayerBuilder) LOAD(clip string, parameters *map[string]string) (*Response, error) {
	cmd := types.CommandLoad{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Clip:       clip,
		Parameters: parameters,
	}
	return b.client.Send(cmd)
}

// PLAY plays content on the layer
func (b *LayerBuilder) PLAY(clip *string, parameters *map[string]string) (*Response, error) {
	cmd := types.CommandPlay{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Clip:       clip,
		Parameters: parameters,
	}
	return b.client.Send(cmd)
}

// PAUSE pauses playback on the layer
func (b *LayerBuilder) PAUSE() (*Response, error) {
	cmd := types.CommandPause{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	return b.client.Send(cmd)
}

// RESUME resumes playback on the layer
func (b *LayerBuilder) RESUME() (*Response, error) {
	cmd := types.CommandResume{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	return b.client.Send(cmd)
}

// STOP stops playback on the layer
func (b *LayerBuilder) STOP() (*Response, error) {
	cmd := types.CommandStop{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	return b.client.Send(cmd)
}

// CLEAR clears the layer
func (b *LayerBuilder) CLEAR() (*Response, error) {
	cmd := types.CommandClear{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
	}
	return b.client.Send(cmd)
}

// CALL calls a function on the layer
func (b *LayerBuilder) CALL(params map[string]string) (*Response, error) {
	cmd := types.CommandCall{
		BasicCommand: types.BasicCommand{
			VideoChannel: b.videoChannel,
			Layer:        b.layer,
		},
		Params: params,
	}
	return b.client.Send(cmd)
}

// Direct command methods on Client for commands that don't require a builder

// SWAP swaps layers between channels
func (c *Client) SWAP(channel1, channel2 int, layer1, layer2 *int, transform bool) (*Response, error) {
	cmd := types.CommandSwap{
		VideoChannel1: channel1,
		Layer1:        layer1,
		VideoChannel2: channel2,
		Layer2:        layer2,
		Transform:     transform,
	}
	return c.Send(cmd)
}

// ADD adds a consumer to the specified video channel
func (c *Client) ADD(videoChannel int, consumerIdx *int, consumerName string, parameters map[string]string) (*Response, error) {
	cmd := types.CommandAdd{
		VideoChannel: videoChannel,
		ConsumerIdx:  consumerIdx,
		ConsumerName: consumerName,
		Parameters:   parameters,
	}
	return c.Send(cmd)
}

// REMOVE removes a consumer from the specified video channel
func (c *Client) REMOVE(videoChannel int, consumerIdx *int, parameters *map[string]string) (*Response, error) {
	cmd := types.CommandRemove{
		VideoChannel: videoChannel,
		ConsumerIdx:  consumerIdx,
		Parameters:   parameters,
	}
	return c.Send(cmd)
}

// PRINT sends a print command for the specified video channel
func (c *Client) PRINT(videoChannel int) (*Response, error) {
	cmd := types.CommandPrint{
		VideoChannel: videoChannel,
	}
	return c.Send(cmd)
}

// LOGLEVEL sets the log level
func (c *Client) LOGLEVEL(level types.LogLevel) (*Response, error) {
	cmd := types.CommandLogLevel{
		Level: level,
	}
	return c.Send(cmd)
}

// SET changes the value of a channel variable
func (c *Client) SET(videoChannel int, variable types.SetVariable, value string) (*Response, error) {
	cmd := types.CommandSet{
		VideoChannel: videoChannel,
		Variable:     variable,
		Value:        value,
	}
	return c.Send(cmd)
}

// LOCK performs a lock operation on the specified channel
func (c *Client) LOCK(videoChannel int, action types.LockAction, secret *string) (*Response, error) {
	cmd := types.CommandLock{
		VideoChannel: videoChannel,
		Action:       action,
		Secret:       secret,
	}
	return c.Send(cmd)
}

// PING sends a ping command
func (c *Client) PING(token string) (*Response, error) {
	cmd := types.CommandPing{
		Token: token,
	}
	return c.Send(cmd)
}

// BYE closes the connection
func (c *Client) BYE() (*Response, error) {
	cmd := types.CommandBye{}
	return c.Send(cmd)
}

// KILL kills the server
func (c *Client) KILL() (*Response, error) {
	cmd := types.CommandKill{}
	return c.Send(cmd)
}

// RESTART restarts the server
func (c *Client) RESTART() (*Response, error) {
	cmd := types.CommandRestart{}
	return c.Send(cmd)
}

// Query command methods

var reCINF = regexp.MustCompile(`^"?([^"]+)"?\s+(\S+)\s+(\d+)\s+(\d+)\s+(\d+)\s+([\d/]+)$`)

// CINF returns information about a media file
func (c *Client) CINF(filename string) (returns.CINF, *Response, error) {
	cmd := types.QueryCommandCINF{
		Filename: filename,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return returns.CINF{}, nil, err
	}

	matches := reCINF.FindStringSubmatch(resp.Data[0])
	cinf, err := matchesToCINF(matches)
	if err != nil {
		return returns.CINF{}, resp, err
	}

	return cinf, resp, nil
}

// CLS lists media files in the media folder
func (c *Client) CLS(directory *string) ([]returns.CINF, *Response, error) {
	cmd := types.QueryCommandCLS{
		Directory: directory,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return nil, nil, err
	}

	cls := []returns.CINF{}
	for _, file := range resp.Data {
		matches := reCINF.FindStringSubmatch(file)
		cinf, err := matchesToCINF(matches)
		if err != nil {
			return nil, nil, err
		}
		cls = append(cls, cinf)
	}

	return cls, resp, nil
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
func (c *Client) FLS() (*Response, error) {
	cmd := types.QueryCommandFLS{}
	return c.Send(cmd)
}

// TLS lists template files
func (c *Client) TLS(directory *string) ([]string, *Response, error) {
	cmd := types.QueryCommandTLS{
		Directory: directory,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return nil, nil, err
	}
	return resp.Data, resp, nil
}

// VERSION returns the version of specified component
func (c *Client) VERSION() (string, *Response, error) {
	return c.version("")
}

func (c *Client) VERSIONSERVER() (string, *Response, error) {
	return c.version(types.VersionInfoServer)
}

func (c *Client) VERSIONFLASH() (string, *Response, error) {
	return c.version(types.VersionInfoFlash)
}

func (c *Client) VERSIONTEMPLATEHOST() (string, *Response, error) {
	return c.version(types.VersionInfoTemplateHost)
}

func (c *Client) VERSIONCEF() (string, *Response, error) {
	return c.version(types.VersionInfoCEF)
}

func (c *Client) version(component types.VersionInfo) (string, *Response, error) {
	cmd := types.QueryCommandVersion{
		Component: component,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return "", nil, err
	}
	if len(resp.Data) == 0 {
		return "", resp, nil
	}
	return strings.TrimSpace(resp.Data[0]), resp, nil
}

func (c *Client) INFO() (*Response, error) {
	resp, _, err := c.info("")
	return resp, err
}

func (c *Client) INFOCONFIG() (*Response, returns.CasparConfig, error) {
	resp, data, err := c.info(types.InfoComponentConfig)
	if data != nil {
		config, ok := data.(returns.CasparConfig)
		if !ok {
			return nil, returns.CasparConfig{}, fmt.Errorf("unexpected data type for config info: %T", data)
		}
		return resp, config, nil
	}
	return resp, returns.CasparConfig{}, err
}

func (c *Client) INFOPATHS() (*Response, returns.Paths, error) {
	resp, data, err := c.info(types.InfoComponentPaths)
	if data != nil {
		paths, ok := data.(returns.Paths)
		if !ok {
			return nil, returns.Paths{}, fmt.Errorf("unexpected data type for paths info: %T", data)
		}
		return resp, paths, nil
	}
	return resp, returns.Paths{}, err
}

func (c *Client) INFOSYSTEM() (*Response, returns.SystemInfo, error) {
	resp, data, err := c.info(types.InfoComponentSystem)
	if data != nil {
		systemInfo, ok := data.(returns.SystemInfo)
		if !ok {
			return nil, returns.SystemInfo{}, fmt.Errorf("unexpected data type for system info: %T", data)
		}
		return resp, systemInfo, nil
	}
	return resp, returns.SystemInfo{}, err
}

func (c *Client) INFOSERVER() (*Response, returns.SystemInfo, error) {
	resp, data, err := c.info(types.InfoComponentServer)
	if data != nil {
		systemInfo, ok := data.(returns.SystemInfo)
		if !ok {
			return nil, returns.SystemInfo{}, fmt.Errorf("unexpected data type for server info: %T", data)
		}
		return resp, systemInfo, nil
	}
	return resp, returns.SystemInfo{}, err
}

func (c *Client) INFOQUEUES() (*Response, returns.SystemInfo, error) {
	resp, data, err := c.info(types.InfoComponentQueues)
	if data != nil {
		systemInfo, ok := data.(returns.SystemInfo)
		if !ok {
			return nil, returns.SystemInfo{}, fmt.Errorf("unexpected data type for queues info: %T", data)
		}
		return resp, systemInfo, nil
	}
	return resp, returns.SystemInfo{}, err
}

func (c *Client) INFOTHREADS() (*Response, returns.SystemInfo, error) {
	resp, data, err := c.info(types.InfoComponentThreads)
	if data != nil {
		systemInfo, ok := data.(returns.SystemInfo)
		if !ok {
			return nil, returns.SystemInfo{}, fmt.Errorf("unexpected data type for threads info: %T", data)
		}
		return resp, systemInfo, nil
	}
	return resp, returns.SystemInfo{}, err
}

func (c *Client) info(component types.InfoComponent) (*Response, any, error) {
	cmd := types.QueryCommandInfo{
		Component: component,
	}
	resp, err := c.Send(cmd)
	if err != nil {
		return nil, nil, err
	}

	fullXml := strings.Join(resp.Data, "\n")
	switch component {
	case types.InfoComponentConfig:
		var config returns.CasparConfig
		err := xml.Unmarshal([]byte(fullXml), &config)
		if err != nil {
			return nil, nil, err
		}
		return resp, config, nil

	case types.InfoComponentPaths:
		var paths returns.Paths
		err := xml.Unmarshal([]byte(fullXml), &paths)
		if err != nil {
			return nil, nil, err
		}
		return resp, paths, nil

	case types.InfoComponentSystem, types.InfoComponentServer, types.InfoComponentQueues, types.InfoComponentThreads:
		parts := strings.Split(fullXml, " ")
		if len(parts) != 3 {
			return nil, nil, fmt.Errorf("unexpected format for '%s' info: %s", component, fullXml)
		}

		videoChannel, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid video channel in '%s' info: %s", component, parts[0])
		}

		systemInfo := returns.SystemInfo{
			VideoChannel: videoChannel,
			Mode:         types.VideoMode(parts[1]),
			Status:       parts[2],
		}
		return resp, systemInfo, nil

	default:
		return resp, fullXml, nil
	}
}

// INFOCHANNEL gets information about a channel or layer
func (c *Client) INFOCHANNEL(videoChannel int, layer *int) (*Response, error) {
	cmd := types.QueryCommandInfoChannel{
		VideoChannel: videoChannel,
		Layer:        layer,
	}
	return c.Send(cmd)
}

// INFOTEMPLATE gets information about the specified template
func (c *Client) INFOTEMPLATE(template string) (*Response, error) {
	cmd := types.QueryCommandInfoTemplate{
		Template: template,
	}
	return c.Send(cmd)
}

// INFODELAY gets delay information
func (c *Client) INFODELAY(videoChannel int, layer *int) (*Response, error) {
	cmd := types.QueryCommandInfoDelay{
		VideoChannel: videoChannel,
		Layer:        layer,
	}
	return c.Send(cmd)
}
