package main

import (
	"context"
	"fmt"
	"time"

	"github.com/overlayfox/casparcg-amcp-go"
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

	brightness, err := client.Mixer().Channel(1).Layer(1).GetBrightness()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Brightness: %f\n", brightness)

	err = client.Mixer().Channel(1).Layer(1).SetBrightness(2)
	if err != nil {
		panic(err)
	}

	brightness, err = client.Mixer().Channel(1).Layer(1).GetBrightness()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Brightness: %f\n", brightness)

	time.Sleep(1 * time.Second)

	err = client.Mixer().Channel(1).Layer(1).SetBrightness(1.0)
	if err != nil {
		panic(err)
	}

	brightness, err = client.Mixer().Channel(1).Layer(1).GetBrightness()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Brightness: %f\n", brightness)
}
