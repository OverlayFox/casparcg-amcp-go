package main

import (
	"context"
	"encoding/json"
	"errors"
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

	// Ping server to see if it's alive and responding
	pingResp, err := client.Ping(ptr("CasparCG AMCP Go Client - Info Example\n\n\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping response:", pingResp)

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
	fmt.Println("Generic Info:")
	fmt.Println(string(jsonData) + "\n\n\n")

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
	fmt.Println("Deprecated Queues Info:")
	fmt.Println(string(jsonData) + "\n\n\n")

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
	fmt.Println("Config Info:")
	fmt.Println(string(jsonData) + "\n\n\n")

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
	fmt.Println("Layer Channel Info:")
	fmt.Println(string(jsonData) + "\n\n\n")

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
	fmt.Println("Layer Channel Layer Info:")
	fmt.Println(string(jsonData) + "\n\n\n")

	fmt.Println("Done.")
}

func ptr[T any](v T) *T {
	return &v
}
