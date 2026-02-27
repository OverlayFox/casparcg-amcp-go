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
	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}
	}()

	data, _, err := client.VERSION()
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
