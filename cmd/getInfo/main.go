package main

import (
	"encoding/json"
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

	comp := types.InfoComponentConfig
	_, data, err := client.INFO(&comp)
	if err != nil {
		panic(err)
	}
	formatted := data.(types.CasparConfig)
	fmt.Printf("%v", formatted)

	// Pretty print as JSON
	jsonData, err := json.MarshalIndent(formatted, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
