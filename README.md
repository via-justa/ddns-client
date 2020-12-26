![Go Test](https://github.com/via-justa/ddns-client/workflows/Tests/badge.svg)  ![Language](https://img.shields.io/badge/Language-go-green)   [![Go Report Card](https://goreportcard.com/badge/github.com/via-justa/ddns-client)](https://goreportcard.com/report/github.com/via-justa/ddns-client)  [![license](https://img.shields.io/badge/license-CC-blue)](https://creativecommons.org/licenses/by-nc-sa/4.0/)
# ddns-client

General
-------

ddns-client is a Go client to update public DNS with your current IP (a.k.a DDNS).

Supported Operating systems
---------------------------

The binaries are built for Linux (amd and arm), Windows (i386) and MacOs (darwin i386).
The client should work on any operating system when built from source

Installation
------------

-   Download the relevant version from the [release page](https://github.com/via-justa/ddns-client/releases)
-   Set the relevant Environment variables to your provider.

Running in Docker
------------

```
docker pull viajusta/ddns-client
docker run --rm -e <provider env var>=<val> viajusta/ddns-client -p <provider> -d <comma separated list of FQDN> -i <interval>
```

Build from source
------------

```
git clone git@github.com:via-justa/ddns-client.git
cd ddns-client
go build
```

Command line flags
------------------
-   **-d, --dns**       comma separated list of FQDN of records to set
-   **-p, --provider**  provider hosting the DNS zone
-   **-i, --interval**  interval to check records status in time format (30m, 1h, 2d, 1w ...)
-   **-l, --log**       set log file path to use. default: none (print to console)
-   **-h, --help**      print available options
-   **-v, --version**   print away client version

Supported providers and required environment variables
-----------

|Provider|Environment variables|
|---|---|
| hetzner | `HETZNER_API_KEY` |

Issues and feature requests
-----------

I'm more than happy to reply to any issue of feature request via the github issue tracker.
When opening an issue, please provide the version you're using and any useful information that can be used to investigate the issue.

If you like to suggest improvements please provide relevant use case it solves.

Requests to add providers should include link to the provider API documentation and information if the provider can set a send box environment.

Pull requests are welcomed, just try to keep the structure as close to the existing ones.

License
-----------

<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.
