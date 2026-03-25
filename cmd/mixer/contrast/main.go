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

	contrast, err := client.Mixer(1, 1).GetContrast()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contrast: %f\n", contrast)

	err = client.Mixer(1, 1).SetContrast(0)
	if err != nil {
		panic(err)
	}

	contrast, err = client.Mixer(1, 1).GetContrast()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contrast: %f\n", contrast)

	time.Sleep(1 * time.Second)

	err = client.Mixer(1, 1).SetContrast(1.0)
	if err != nil {
		panic(err)
	}

	contrast, err = client.Mixer(1, 1).GetContrast()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Contrast: %f\n", contrast)
}
