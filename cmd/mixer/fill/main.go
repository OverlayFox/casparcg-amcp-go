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

	originalFill, err := client.Mixer().Channel(1).Layer(1).GetFill()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fill: %+v\n", originalFill)

	newFill := types.MixerParamsFill{
		X:      0.5,
		Y:      0.5,
		XScale: 1.2,
		YScale: 1.2,
	}
	err = client.Mixer().Channel(1).Layer(1).SetFill(newFill)

	if err != nil {
		panic(err)
	}

	fill, err := client.Mixer().Channel(1).Layer(1).GetFill()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fill: %+v\n", fill)

	time.Sleep(2 * time.Second)

	err = client.Mixer().Channel(1).Layer(1).SetFill(types.MixerParamsFill{
		X:      originalFill.X,
		Y:      originalFill.Y,
		XScale: originalFill.XScale,
		YScale: originalFill.YScale,
	})
	if err != nil {
		panic(err)
	}

	fill, err = client.Mixer().Channel(1).Layer(1).GetFill()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fill: %+v\n", fill)
}
