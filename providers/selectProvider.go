package provider

import (
	"fmt"

	"github.com/via-justa/ddns-client/providers/hetzner"
	"github.com/via-justa/ddns-client/providers/namecheap"
)

// Provider enables implementing a custom DNS provider
type Provider interface {
	Update(host, domain string, currentIP string) error
}

// NewDNSProvider accept provider name and return Provider.
// If provider name does not exists or the provider initialization failed return an error
func NewDNSProvider(provider string) (Provider, error) {
	// nolint:gocritic
	switch provider {
	case "hetzner":
		return hetzner.NewDNSProvider()
	case "namecheap":
		return namecheap.NewDNSProvider()
	}

	return nil, fmt.Errorf("No valid proveder has been selected. Please use valid provider")
}
