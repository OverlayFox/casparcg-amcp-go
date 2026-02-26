package main

import (
	"fmt"

	"github.com/overlayfox/casparcg-amcp-go"
)

func main() {
	client := casparcg.NewClient("127.0.0.1", 5250)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_, data, err := client.INFOCONFIG()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", data)
}
