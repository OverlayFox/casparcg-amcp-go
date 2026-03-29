package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/overlayfox/casparcg-amcp-go"
	"github.com/overlayfox/casparcg-amcp-go/types"
)

// main demonstrates how to play a video clip on a specific channel and layer, then pause, resume, and stop it using the CasparCG AMCP Go client.
//
// This is not best practice for a real application - it's just a simple example to show how to use the client.
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

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("  CasparCG AMCP Go Client - Video Playback Example")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	// Ping server to see if it's alive and responding
	resp, err := client.Ping(ptr("CasparCG AMCP Go Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Ping Response")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(resp)
	fmt.Println()

	// Play a clip on channel 1, layer 10 without any optional parameters for 2 seconds
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Playing clip on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))
	err = client.Layer().Channel(1).Layer(10).Play(types.LayerPlay{ClipName: ptr("BACKGROUNDLOOP")})
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	fmt.Println("OK")
	fmt.Println()
	time.Sleep(2 * time.Second)

	// Pause the clip for 1 second
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Pausing clip on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))
	err = client.Layer().Channel(1).Layer(10).Pause()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	fmt.Println("OK")
	fmt.Println()
	time.Sleep(1 * time.Second)

	// Resume the clip for 2 seconds
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Resuming clip on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))
	err = client.Layer().Channel(1).Layer(10).Resume()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	fmt.Println("OK")
	fmt.Println()
	time.Sleep(2 * time.Second)

	// Stop the clip on channel 1, layer 10
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Stopping clip on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))
	err = client.Layer().Channel(1).Layer(10).Stop()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	fmt.Println("OK")
	fmt.Println()

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Done")
	fmt.Println(strings.Repeat("=", 80))
}

func ptr[T any](v T) *T {
	return &v
}
