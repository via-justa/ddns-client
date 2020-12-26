package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
	ipcheck "github.com/via-justa/ddns-client/IPCheck"
	provider "github.com/via-justa/ddns-client/providers"
)

var (
	appVersion = "undefined"
	prov       string
	dnsName    []string
	interval   string
	logfile    string
	help       bool
	version    bool
)

func init() {
	pflag.StringVarP(&prov, "provider", "p", "hetzner", "provider hosting the DNS zone")
	pflag.StringSliceVarP(&dnsName, "dns", "d", nil, "comma separated list of FQDN of records to set")
	pflag.StringVarP(&interval, "interval", "i", "", "Interval to check records status (e.g. 30m, 1h)")
	pflag.StringVarP(&logfile, "log", "l", "", "set log file path to use. default: none (print to console)")
	pflag.BoolVarP(&help, "help", "h", false, "print available options")
	pflag.BoolVarP(&version, "version", "v", false, "print ddns-client version")
	pflag.Parse()
}

func main() {
	if help {
		pflag.PrintDefaults()
		return
	}

	if version {
		log.Printf("ddns-client version: %v", appVersion)
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

	if interval != "" {
		for {
			err := updateRecords()
			if err != nil {
				log.Fatal(err)
			}

			t, err := time.ParseDuration(interval)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(t)
		}
	} else {
		err := updateRecords()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func updateRecords() error {
	ip, err := ipcheck.GetCurrentIP()
	if err != nil {
		return err
	}

	log.Printf("current IP: %v", ip)

	prov, err := provider.NewDNSProvider(prov)
	if err != nil {
		return err
	}

	for i := range dnsName {
		log.Printf("DNS Name: %v", dnsName[i])
		idx := strings.Index(dnsName[i], ".")

		err = prov.Update(dnsName[i][:idx], dnsName[i][idx+1:], ip)
		if err != nil {
			return err
		}
	}

	return nil
}
