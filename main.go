package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	ipcheck "github.com/via-justa/ddns-client/IPCheck"
	provider "github.com/via-justa/ddns-client/providers"
)

var (
	appVersion = "undefined"
	prov       string
	dnsName    string
	logfile    string
	help       bool
	version    bool
)

func init() {
	pflag.StringVarP(&prov, "provider", "p", "hetzner", "provider hosting the DNS zone")
	pflag.StringVarP(&dnsName, "dns", "d", "", "FQDN of record to set")
	pflag.StringVarP(&logfile, "log", "l", "", "set log file path to use. default: none (print to console)")
	pflag.BoolVarP(&help, "help", "h", false, "print available options")
	pflag.BoolVarP(&version, "version", "v", false, "print away-client version")
	pflag.Parse()
}

func main() {
	if help {
		pflag.PrintDefaults()
		return
	}

	if version {
		log.Printf("away-client version: %v", appVersion)
		return
	}

	// write log to file in addition to writing to console
	if len(logfile) > 0 {
		var f *os.File

		f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		// nolint:errcheck,gosec
		defer f.Close()

		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)
	}

	ip, err := ipcheck.GetCurrentIP()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("current IP: %v", ip)

	prov, err := provider.NewDNSProvider(prov)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DNS Name: %v", dnsName)
	idx := strings.Index(dnsName, ".")

	err = prov.Update(dnsName[:idx], dnsName[idx+1:], ip)
	if err != nil {
		log.Fatal(err)
	}
}
