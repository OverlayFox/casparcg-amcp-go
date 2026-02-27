package main

import (
	"time"

	"github.com/overlayfox/casparcg-amcp-go"
)

func main() {
	client := casparcg.NewClient("127.0.0.1", 5250)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}
	}()

	clip := "BACKGROUNDLOOP"
	resp, err := client.Layer(1, 10).PLAY(&clip, nil)
	if err != nil {
		panic(err)
	}
	println("CG ADD response:", resp.Code, resp.Message)

	time.Sleep(2 * time.Second)

	resp, err = client.Layer(1, 10).STOP()
	if err != nil {
		panic(err)
	}
	println("CG STOP response:", resp.Code, resp.Message)
}
