package main

import (
	"context"
	"errors"
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

	clip := "BACKGROUNDLOOP"
	err = client.Layer(1, 10).Play(&clip, nil)
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
	time.Sleep(2 * time.Second)

	err = client.Layer(1, 10).Stop()
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}
}
