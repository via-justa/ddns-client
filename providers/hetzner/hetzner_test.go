// Package hetzner implements a DNS provider for dynamic DNS update.
package hetzner

import (
	"testing"
)

func TestNewDNSProvider(t *testing.T) {
	t.Run("working config", func(t *testing.T) {
		_, err := NewDNSProvider()
		if err != nil {
			t.Errorf("NewDNSProvider() error = %v", err)
			return
		}
	})
}
