// Package namecheap implements a DNS provider for dynamic DNS update.
package namecheap

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/via-justa/ddns-client/providers/namecheap/internal"
)

// Environment variables names.
const envNamespace = "NAMECHEAP_"

// Config is used to configure the creation of the DNSProvider.
type Config struct{}

// NewDefaultConfig returns a default configuration for the DNSProvider.
func NewDefaultConfig() *Config {
	return &Config{}
}

// DNSProvider implements the dns.Provider interface.
type DNSProvider struct {
	config *Config
	client *internal.Client
}

// NewDNSProvider returns a DNSProvider instance configured for namecheap.
// Credentials must be passed in the environment variable: NAMECHEAP_API_KEY.
func NewDNSProvider() (*DNSProvider, error) {
	config := NewDefaultConfig()

	return NewDNSProviderConfig(config)
}

// NewDNSProviderConfig return a DNSProvider instance configured for namecheap.
func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	if config == nil {
		return nil, errors.New("namecheap: the configuration of the DNS provider is nil")
	}

	client := internal.NewClient()

	return &DNSProvider{config: config, client: client}, nil
}

// Update creates or update the requested record.
func (dp *DNSProvider) Update(host, domain string, currentIP string) error {
	// check the domain password env var exists
	envVar := strings.ToUpper(envNamespace +
		strings.ReplaceAll(strings.ReplaceAll(domain, "-", "_"), ".", "_") +
		"_PASSWORD")

	ddnsDomainPassword := os.Getenv(envVar)
	if ddnsDomainPassword == "" {
		return errors.New("namecheap: some credentials information are missing: " + envVar)
	}

	requestParam := internal.RequestParams{
		Host:     host,
		Domain:   domain,
		IP:       currentIP,
		Password: ddnsDomainPassword,
	}

	err := dp.client.UpdateRecord(&requestParam)
	if err != nil {
		return fmt.Errorf("namecheap: failed to update record: %s.%s: %w", host, domain, err)
	}

	return nil
}
