package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/overlayfox/casparcg-amcp-go"
)

// main demonstrates how to query various types of information from the CasparCG server using the CasparCG AMCP Go client.
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
	fmt.Println("CasparCG AMCP Go Client - Info Example")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	// Ping server to see if it's alive and responding
	pingResp, err := client.Ping(ptr("CasparCG AMCP Go Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Ping Response")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(pingResp)
	fmt.Println()

	// Get generic info about the server and print it as JSON
	resp, err := client.Query().Info().Generic()
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
	fmt.Println("Server Generic Info")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(string(jsonData))
	fmt.Println()

	// Some functions are marked as "deprecated" but they still work.
	// They just won't return sensible data.
	resp, err = client.Query().Info().Queues()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	jsonData, err = json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Queues Info (Deprecated)")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(string(jsonData))
	fmt.Println()

	// Get config info about the server and print it as JSON
	configResp, err := client.Query().Info().Config()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	jsonData, err = json.MarshalIndent(configResp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Server Config Info")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(string(jsonData))
	fmt.Println()

	// Some query commands were moved to different builders to better reflect their purpose.
	layerResp, err := client.Layer().Channel(1).Info().Generic()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	jsonData, err = json.MarshalIndent(layerResp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Channel 1 Info")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(string(jsonData))
	fmt.Println()

	// Some query commands were moved to different builders to better reflect their purpose.
	layerResp, err = client.Layer().Channel(1).Layer(10).Info().Generic()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	jsonData, err = json.MarshalIndent(layerResp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Channel 1, Layer 10 Info")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(string(jsonData))
	fmt.Println()

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Done")
	fmt.Println(strings.Repeat("=", 80))
}

func ptr[T any](v T) *T {
	return &v
}
