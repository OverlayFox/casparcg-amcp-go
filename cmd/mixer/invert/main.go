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

	invertState, err := client.Mixer(1, 1).GetInvert()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invert state: %t\n", invertState)

	err = client.Mixer(1, 1).SetInvert(true)
	if err != nil {
		panic(err)
	}

	invertState, err = client.Mixer(1, 1).GetInvert()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invert state: %t\n", invertState)

	time.Sleep(1 * time.Second)

	err = client.Mixer(1, 1).SetInvert(false)
	if err != nil {
		panic(err)
	}

	invertState, err = client.Mixer(1, 1).GetInvert()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invert state: %t\n", invertState)
}
