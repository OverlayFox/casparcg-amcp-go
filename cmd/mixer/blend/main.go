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

	mode, err := client.Mixer(1, 1).GetBlendMode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Blend mode: %s\n", mode)

	err = client.Mixer(1, 1).SetBlendMode(types.BlendModeScreen)
	if err != nil {
		panic(err)
	}

	mode, err = client.Mixer(1, 1).GetBlendMode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Blend mode: %s\n", mode)

	time.Sleep(1 * time.Second)

	err = client.Mixer(1, 1).SetBlendMode(types.BlendModeNormal)
	if err != nil {
		panic(err)
	}

	mode, err = client.Mixer(1, 1).GetBlendMode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Blend mode: %s\n", mode)
}
