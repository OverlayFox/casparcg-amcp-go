package responses

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

type QueryChannelInfo struct {
	ChannelIndex int
	VideoMode    types.VideoMode
	Status       string
}

func PartsToQueryChannelInfo(parts []string) (QueryChannelInfo, error) {
	if len(parts) != 3 {
		return QueryChannelInfo{}, fmt.Errorf("unexpected format for channel info: %s", strings.Join(parts, " "))
	}

	videoChannel, err := strconv.Atoi(parts[0])
	if err != nil {
		return QueryChannelInfo{}, fmt.Errorf("invalid video channel in channel info: %s", parts[0])
	}

	channelInfo := QueryChannelInfo{
		ChannelIndex: videoChannel,
		VideoMode:    types.VideoMode(parts[1]),
		Status:       parts[2],
	}
	return channelInfo, nil
}

func ResponseToQueryChannelInfo(response []string) ([]QueryChannelInfo, error) {
	channelInfo := make([]QueryChannelInfo, 0, len(response))
	for _, line := range response {
		parts := strings.Split(line, " ")
		info, err := PartsToQueryChannelInfo(parts)
		if err != nil {
			return nil, err
		}
		channelInfo = append(channelInfo, info)
	}
	return channelInfo, nil
}
