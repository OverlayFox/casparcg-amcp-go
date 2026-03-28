package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/overlayfox/casparcg-amcp-go"
	"github.com/overlayfox/casparcg-amcp-go/types"
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

	err = client.CG().Channel(1).Layer(1).CGLayer(1).Add(types.CGAdd{Template: "L3", PlayOnLoad: true, Data: jsonString})
	if err != nil {
		var casparErr casparcg.CasparCGError
		if errors.As(err, &casparErr) {
			fmt.Printf("CasparCG error: %d - %s\n", casparErr.Code, casparErr.Message)
		}
		panic(err)
	}

	// jsonData, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(jsonData))
}
