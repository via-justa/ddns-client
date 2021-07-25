package internal

import "encoding/xml"

// RequestParams the required parameters for namecheap dynamic DNS update request
type RequestParams struct {
	Host     string
	Domain   string
	Password string
	IP       string
}

// Response represent namecheap update request response
type Response struct {
	XMLName  xml.Name `xml:"interface-response"`
	IP       string   `xml:"IP"`
	ErrCount int      `xml:"ErrCount"`
	Errors   Errors   `xml:"errors"`
	Done     bool     `xml:"Done"`
}

// Errors represent namecheap update request response error string
type Errors struct {
	Error string `xml:"Err1"`
}
