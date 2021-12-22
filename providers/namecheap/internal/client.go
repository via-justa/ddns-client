package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// defaultBaseURL represents the API endpoint to call.
const defaultBaseURL = "https://dynamicdns.park-your-domain.com/update"

// Client the NameCheap client.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient Creates a new NameCheap client.
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    defaultBaseURL,
	}
}

func identReader(encoding string, input io.Reader) (io.Reader, error) {
	return input, nil
}

/*
UpdateRecord update a Dynamic DNS record.
https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-do-i-use-a-browser-to-dynamically-update-the-hosts-ip/
*/ // nolint: lll
func (c *Client) UpdateRecord(requestParams *RequestParams) error {
	resp, err := c.do(requestParams)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("namecheap: server returned error code %v", resp.StatusCode)
	}

	var responseBody Response

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = identReader
	if err = decoder.Decode(&responseBody); err != nil {
		return fmt.Errorf("namecheap: %w", err)
	}

	if responseBody.ErrCount != 0 {
		return fmt.Errorf(responseBody.Errors.Error)
	}

	return nil
}

func (c *Client) do(requestParams *RequestParams) (*http.Response, error) {
	param := url.Values{}
	param.Set("host", requestParams.Host)
	param.Set("domain", requestParams.Domain)
	param.Set("password", requestParams.Password)
	param.Set("ip", requestParams.IP)

	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"?"+param.Encode(), nil) // nolint: noctx
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml")

	return c.HTTPClient.Do(req)
}
