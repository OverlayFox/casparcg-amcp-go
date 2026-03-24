package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/overlayfox/casparcg-amcp-go"
	"github.com/overlayfox/casparcg-amcp-go/types/returns"
)

func main() {
	client := casparcg.NewClient("127.0.0.1", 5250)
	err := client.Connect(context.TODO()) // replace with a proper context when implementing this
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}
	}()

	info, err := client.Mixer(1, 1).ChromaInfo()
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	err = client.Mixer(1, 1).ChromaEnable(returns.MixerChromaInfo{
		Enabled:                 true,
		TargetHue:               120,
		HueWidth:                0.1,
		MinSaturation:           0,
		MinBrightness:           0,
		Softness:                0.1,
		SpillSuppress:           0.1,
		SpillSuppressSaturation: 0.7,
		ShowMask:                false,
	}, &casparcg.ChromaFade{
		Duration: 25, // in frames
		Tween:    "linear",
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(600 * time.Millisecond)
	info, err = client.Mixer(1, 1).ChromaInfo()
	if err != nil {
		panic(err)
	}
	jsonData, err = json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	err = client.Mixer(1, 1).ChromaDisable(nil)
	if err != nil {
		panic(err)
	}

	info, err = client.Mixer(1, 1).ChromaInfo()
	if err != nil {
		panic(err)
	}
	jsonData, err = json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
