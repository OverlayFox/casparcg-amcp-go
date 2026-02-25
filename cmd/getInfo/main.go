package main

import (
	"fmt"

	"github.com/overlayfox/casparcg-amcp-go"
	"github.com/overlayfox/casparcg-amcp-go/types"
)

func main() {
	client := casparcg.NewClient("127.0.0.1", 5250)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	comp := types.InfoComponentPaths
	resp, _, err := client.INFO(&comp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resp)
}
