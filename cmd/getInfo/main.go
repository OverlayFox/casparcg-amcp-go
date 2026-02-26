package main

import (
	"encoding/json"
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

	resp, data, err := client.INFOTHREADS()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resp.Data)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", jsonData)
}
