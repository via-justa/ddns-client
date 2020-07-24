// Package hetzner implements a DNS provider for dynamic DNS update.
package hetzner

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/via-justa/ddns-client/providers/hetzner/internal"
)

// Environment variables names.
const (
	envNamespace = "HETZNER_"
	envAPIKey    = envNamespace + "API_KEY"
)

// Config is used to configure the creation of the DNSProvider.
type Config struct {
	APIKey     string
	HTTPClient *http.Client
}

// NewDefaultConfig returns a default configuration for the DNSProvider.
func NewDefaultConfig() *Config {
	return &Config{
		HTTPClient: &http.Client{},
	}
}

// DNSProvider implements the dns.Provider interface.
type DNSProvider struct {
	config *Config
	client *internal.Client
}

// NewDNSProvider returns a DNSProvider instance configured for hetzner.
// Credentials must be passed in the environment variable: HETZNER_API_KEY.
func NewDNSProvider() (*DNSProvider, error) {
	config := NewDefaultConfig()
	config.APIKey = os.Getenv(envAPIKey)

	return NewDNSProviderConfig(config)
}

// NewDNSProviderConfig return a DNSProvider instance configured for hetzner.
func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	if config == nil {
		return nil, errors.New("hetzner: the configuration of the DNS provider is nil")
	}

	if config.APIKey == "" {
		return nil, errors.New("hetzner: some credentials information are missing: HETZNER_API_KEY")
	}

	client := internal.NewClient(config.APIKey)

	if config.HTTPClient != nil {
		client.HTTPClient = config.HTTPClient
	}

	return &DNSProvider{config: config, client: client}, nil
}

// Update creates or update the requested record.
func (dp *DNSProvider) Update(host, domain string, currentIP string) error {
	zoneID, err := dp.client.GetZoneID(domain)
	if err != nil {
		return fmt.Errorf("hetzner: failed to find zone: domain=%s: %w", domain, err)
	}

	existingRecord, err := dp.client.GetRecord(host, "A", zoneID)
	if err != nil {
		log.Printf("hetzner: %v, creating new one", err)

		newRecord := internal.DNSRecord{
			Name:   host,
			Type:   "A",
			Value:  currentIP,
			TTL:    600,
			ZoneID: zoneID,
		}

		err = dp.client.CreateRecord(&newRecord)
		if err != nil {
			return fmt.Errorf("hetzner: failed to create record: domain=%s host=%s: %w", domain, host, err)
		}

		return nil
	}

	if existingRecord != nil && existingRecord.Value == currentIP {
		log.Printf("hetzner: Existing record value and current IP are the same, nothing to do.")
		return nil
	}

	log.Printf("hetzner: Existing record value and current IP differed, updating existing record.")

	err = dp.client.UpdateRecord(currentIP, existingRecord)
	if err != nil {
		return fmt.Errorf("hetzner: failed to update record: domain=%s host=%s: %w", domain, host, err)
	}

	return nil
}
