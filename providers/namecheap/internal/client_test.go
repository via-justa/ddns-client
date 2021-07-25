package internal

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_UpdateRecord(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/update", func(rw http.ResponseWriter, req *http.Request) {
		file, err := os.Open("./fixtures/create_record.xml")
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

	client := NewClient()
	client.BaseURL = server.URL + "/update"

	requestParam := RequestParams{
		Host:     "test",
		Domain:   "example.com",
		IP:       "99.99.99.99",
		Password: "dummyPassword",
	}

	err := client.UpdateRecord(&requestParam)
	require.NoError(t, err)
}

func TestClient_UpdateRecordError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/update", func(rw http.ResponseWriter, req *http.Request) {
		file, err := os.Open("./fixtures/create_record_error.xml")
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

	client := NewClient()
	client.BaseURL = server.URL + "/update"

	requestParam := RequestParams{
		Host:     "test",
		Domain:   "example.com",
		IP:       "99.99.99.99",
		Password: "dummyPassword",
	}

	err := client.UpdateRecord(&requestParam)
	require.Error(t, err)

	if err.Error() != "Domain name not found" {
		t.Errorf("not the expected error: %v", err.Error())
	}
}
