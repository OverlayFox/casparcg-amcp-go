//go:build integration

package casparcg_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/overlayfox/casparcg-amcp-go"
)

var client *casparcg.Client

func TestMain(m *testing.M) {
	// Check if CasparCG Server is running before executing tests
	client = casparcg.NewClient("localhost", 5250)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to CasparCG Server: %v", err)
	}

	os.Exit(m.Run())
}

func TestIntegration(t *testing.T) {
	t.Run("PING", func(t *testing.T) {
		resp, err := client.Ping("")
		if err != nil {
			t.Fatalf("Failed to ping CasparCG Server: %v", err)
		}
		if resp != "PONG" {
			t.Fatalf("Unexpected response from CasparCG Server: %v", resp)
		}
	})
}
