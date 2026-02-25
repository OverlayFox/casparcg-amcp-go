package main

import "github.com/overlayfox/casparcg-amcp-go"

// This examples demonstrates how to use the PING command to check if the CasparCG server is responsive.
// The PING command can also be used to measure latency by including a message that will be echoed back in the response.
func main() {
	client := casparcg.NewClient("127.0.0.1", 5250)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	resp, err := client.PING("")
	if err != nil {
		panic(err)
	}
	println(resp.Message)
}
