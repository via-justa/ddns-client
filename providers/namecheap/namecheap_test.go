// Package namecheap implements a DNS provider for dynamic DNS update.
package namecheap

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/via-justa/ddns-client/providers/namecheap/internal"
)

func TestNewDNSProvider_missing_domain_env(t *testing.T) {
	provider, err := NewDNSProvider()
	if err != nil {
		t.Errorf("NewDNSProvider() error = %v", err)
		return
	}

	err = provider.Update("test", "example-domain.com", "99.99.99.99")
	require.Error(t, err)

	if err.Error() != "namecheap: some credentials information are missing: NAMECHEAP_EXAMPLE_DOMAIN_COM_PASSWORD" {
		t.Errorf("not the expected error: %v", err.Error())
	}
}

func TestNewDNSProvider(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/update", func(rw http.ResponseWriter, req *http.Request) {
		file, err := os.Open("./internal/fixtures/create_record.xml")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() { _ = file.Close() }()

		_, err = io.Copy(rw, file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	provider, err := NewDNSProvider()
	if err != nil {
		t.Errorf("NewDNSProvider() error = %v", err)
		return
	}

	client := internal.NewClient()
	client.BaseURL = server.URL + "/update"

	provider.client = client

	os.Setenv("NAMECHEAP_EXAMPLE_DOMAIN_COM_PASSWORD", "dummyPassword")

	err = provider.Update("test", "example-domain.com", "99.99.99.99")
	require.NoError(t, err)
}
