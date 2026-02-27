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

	resp, err := client.CG(1, 10).ADD(1, "TITLE", true, nil)
	if err != nil {
		panic(err)
	}
	println("CG ADD response:", resp.Code, resp.Message)

	time.Sleep(2 * time.Second)

	resp, err = client.CG(1, 10).STOP(1)
	if err != nil {
		panic(err)
	}
	println("CG STOP response:", resp.Code, resp.Message)
}
