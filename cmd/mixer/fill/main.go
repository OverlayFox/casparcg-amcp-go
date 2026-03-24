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

	originalFill, err := client.Mixer(1, 1).GetFill()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fill: %+v\n", originalFill)

	newFill := types.MixerParamsFill{
		X:      float32Ptr(0.5),
		Y:      float32Ptr(0.5),
		XScale: float32Ptr(1.2),
		YScale: float32Ptr(1.2),
	}
	err = client.Mixer(1, 1).SetFill(newFill, &types.Fade{Duration: 25, Tween: types.TweenTypeEaseInSine})

	if err != nil {
		panic(err)
	}

	fill, err := client.Mixer(1, 1).GetFill()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fill: %+v\n", fill)

	time.Sleep(2 * time.Second)

	err = client.Mixer(1, 1).SetFill(types.MixerParamsFill{
		X:      &originalFill.X,
		Y:      &originalFill.Y,
		XScale: &originalFill.XScale,
		YScale: &originalFill.YScale,
	}, nil)
	if err != nil {
		panic(err)
	}

	fill, err = client.Mixer(1, 1).GetFill()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fill: %+v\n", fill)
}

func float32Ptr(f float32) *float32 {
	return &f
}
