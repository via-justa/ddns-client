package ipcheck

import (
	"io/ioutil"
	"net"
	"net/http"
)

const url = "http://icanhazip.com"

// GetCurrentIP get the body of request to "http://icanhazip.com" and return IP in response
func GetCurrentIP() (string, error) {
	cli := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}

	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes[:len(bodyBytes)-1])

	ip := net.ParseIP(bodyString)

	return ip.String(), nil
}
