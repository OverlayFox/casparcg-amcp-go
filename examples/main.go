package main

import (
	"fmt"
	"log"

	"github.com/overlayfox/casparcg-amcp-go"
)

func main() {
	// Create a new client
	client := casparcg.NewClient("localhost", 5250)

	// Connect to the server
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	// Example 1: Use CG builder for template commands (as requested)
	// client.CG(1, 12).STOP(2)
	resp, err := client.CG(1, 12).STOP(2)
	if err != nil {
		log.Printf("Failed to stop CG: %v", err)
	} else {
		fmt.Printf("CG STOP response: Code=%d, Message=%s\n", resp.Code, resp.Message)
	}

	// Example 2: Add a template and play it
	templateData := `{"f0":"Hello World"}`
	resp, err = client.CG(1, 10).ADD(1, "lower_third", true, &templateData)
	if err != nil {
		log.Printf("Failed to add template: %v", err)
	} else {
		fmt.Printf("CG ADD response: Code=%d, Message=%s\n", resp.Code, resp.Message)
	}

	// Example 3: Use Layer builder for playback commands
	clip := "AMB"
	resp, err = client.Layer(1, 10).PLAY(&clip, nil)
	if err != nil {
		log.Printf("Failed to play: %v", err)
	} else {
		fmt.Printf("PLAY response: Code=%d, Message=%s\n", resp.Code, resp.Message)
	}

	// Example 4: Load a clip
	resp, err = client.Layer(1, 11).LOAD("myclip", nil)
	if err != nil {
		log.Printf("Failed to load: %v", err)
	} else {
		fmt.Printf("LOAD response: Code=%d, Message=%s\n", resp.Code, resp.Message)
	}

	// Example 5: Query commands
	// resp, err = client.CLS(nil)
	// if err != nil {
	// 	log.Printf("Failed to list media: %v", err)
	// } else {
	// 	fmt.Printf("CLS response: Code=%d, Data lines=%d\n", resp.Code, len(resp.Data))
	// 	for _, line := range resp.Data {
	// 		fmt.Printf("  %s\n", line)
	// 	}
	// }

	// Example 6: Get server version
	// resp, err = client.VERSION(nil)
	// if err != nil {
	// 	log.Printf("Failed to get version: %v", err)
	// } else {
	// 	fmt.Printf("VERSION response: Code=%d, Message=%s\n", resp.Code, resp.Message)
	// }

	// Example 7: Channel operations
	resp, err = client.Layer(1, 10).PAUSE()
	if err != nil {
		log.Printf("Failed to pause: %v", err)
	} else {
		fmt.Printf("PAUSE response: Code=%d\n", resp.Code)
	}

	resp, err = client.Layer(1, 10).RESUME()
	if err != nil {
		log.Printf("Failed to resume: %v", err)
	} else {
		fmt.Printf("RESUME response: Code=%d\n", resp.Code)
	}

	resp, err = client.Layer(1, 10).STOP()
	if err != nil {
		log.Printf("Failed to stop: %v", err)
	} else {
		fmt.Printf("STOP response: Code=%d\n", resp.Code)
	}

	// Example 8: Clear a layer
	resp, err = client.Layer(1, 10).CLEAR()
	if err != nil {
		log.Printf("Failed to clear: %v", err)
	} else {
		fmt.Printf("CLEAR response: Code=%d\n", resp.Code)
	}
}
