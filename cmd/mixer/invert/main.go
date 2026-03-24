package main

import (
	"context"
	"fmt"

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

	invertState, err := client.Mixer(1, 1).GetInvertState()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invert state: %t\n", invertState)

	err = client.Mixer(1, 1).Invert(true)
	if err != nil {
		panic(err)
	}

	invertState, err = client.Mixer(1, 1).GetInvertState()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invert state: %t\n", invertState)

	err = client.Mixer(1, 1).Invert(false)
	if err != nil {
		panic(err)
	}

	invertState, err = client.Mixer(1, 1).GetInvertState()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invert state: %t\n", invertState)
}
