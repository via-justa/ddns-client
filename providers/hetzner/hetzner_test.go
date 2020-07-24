// Package hetzner implements a DNS provider for dynamic DNS update.
package hetzner

import (
	"os"
	"testing"
)

var testEnvVal = "test"

func TestNewDNSProvider(t *testing.T) {
	t.Run("working config", func(t *testing.T) {
		preTestEnvVal := os.Getenv(envAPIKey)
		os.Setenv(envAPIKey, testEnvVal)
		got, err := NewDNSProvider()
		if err != nil {
			t.Errorf("NewDNSProvider() error = %v", err)
			return
		}
		if got.config.APIKey != testEnvVal {
			t.Errorf("NewDNSProvider() error = Env key not match")
			return
		}
		os.Setenv(envAPIKey, preTestEnvVal)
	})
}
