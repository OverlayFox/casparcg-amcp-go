package main

import (
	"context"
	"fmt"
	"time"

	"github.com/overlayfox/casparcg-amcp-go"
	"github.com/overlayfox/casparcg-amcp-go/types"
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

	originalLevels, err := client.Mixer(1, 1).GetLevels()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Levels: %+v\n", originalLevels)

	newLevels := types.MixerInfoLevels{
		MinInput:  0.0627,
		MaxInput:  0.922,
		Gamma:     1,
		MinOutput: 0,
		MaxOutput: 1,
	}
	err = client.Mixer(1, 1).SetLevels(newLevels, &types.Fade{Duration: 25, Tween: types.TweenTypeEaseInSine})
	if err != nil {
		panic(err)
	}

	levels, err := client.Mixer(1, 1).GetLevels()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Levels: %+v\n", levels)

	time.Sleep(2 * time.Second)

	err = client.Mixer(1, 1).SetLevels(originalLevels, nil)
	if err != nil {
		panic(err)
	}

	levels, err = client.Mixer(1, 1).GetLevels()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Levels: %+v\n", levels)
}
