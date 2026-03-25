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

	saturation, err := client.Mixer(1, 1).GetSaturation()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saturation: %f\n", saturation)

	err = client.Mixer(1, 1).SetSaturation(0)
	if err != nil {
		panic(err)
	}

	saturation, err = client.Mixer(1, 1).GetSaturation()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saturation: %f\n", saturation)

	time.Sleep(1 * time.Second)

	err = client.Mixer(1, 1).SetSaturation(1.0)
	if err != nil {
		panic(err)
	}

	saturation, err = client.Mixer(1, 1).GetSaturation()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saturation: %f\n", saturation)
}
