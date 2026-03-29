package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("CasparCG AMCP Go Client - Template & Mixer Example")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	//
	// Ping server to see if it's alive and responding
	//
	pingResp, err := client.Ping(ptr("CasparCG AMCP Go Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Ping Response")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(pingResp)
	fmt.Println()

	//
	// List all templates on the server
	//
	resp, err := client.Query().TLS(nil)
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("TLS Info")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(string(jsonData))
	fmt.Println()

	//
	// Add first available template
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Adding first template on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))

	err = client.CG().Channel(1).Layer(10).CGLayer(0).Add(types.CGAdd{
		Template:   resp[0],
		PlayOnLoad: false,
	})
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	fmt.Println("OK")
	fmt.Println()

	//
	// Get scale and position info about the template we just added
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Scaling template on channel 1, layer 10 to 50%...")
	fmt.Println(strings.Repeat("-", 80))

	posScale, err := client.Mixer().Channel(1).Layer(1).GetFill()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	jsonData, err = json.MarshalIndent(posScale, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println()

	//
	// Set scale to 50%
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Setting scale to 50%...")
	fmt.Println(strings.Repeat("-", 80))

	// NOTE: The parameter for this function is taken from the responses package not from types.
	// This is done to make it easier to modify the clip properties without having to convert between types and responses.
	err = client.Mixer().Channel(1).Layer(10).SetFill(types.MixerParamsFill{
		XScale: 0.5,
		YScale: 0.5,
		X:      0.5,
		Y:      0.5,
	})
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	fmt.Println("OK")
	fmt.Println()

	//
	// Play the template out
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Playing template on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))

	err = client.CG().Channel(1).Layer(10).CGLayer(0).Play()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("OK")
	fmt.Println()

	//
	// Slide the template around
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Sliding template around on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))

	err = client.Mixer().Fade(&types.Fade{
		Duration: 100,
		Tween:    types.TweenTypeEaseInSine}).
		Channel(1).Layer(10).SetFill(types.MixerParamsFill{
		XScale: 0.7,
		YScale: 0.7,
		X:      0.1,
		Y:      0.1,
	})
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	time.Sleep(4 * time.Second)
	fmt.Println("OK")
	fmt.Println()

	//
	// Stop the template
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Stopping template on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))

	err = client.CG().Channel(1).Layer(10).CGLayer(0).Stop()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	time.Sleep(4 * time.Second)
	fmt.Println("OK")
	fmt.Println()

	//
	// Clear all transforms on the layer to leave a clean slate
	//
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Clearing all transforms on channel 1, layer 10...")
	fmt.Println(strings.Repeat("-", 80))

	err = client.Mixer().Channel(1).Layer(10).Clear()
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
