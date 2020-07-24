VERSION=

DOT:= .
DASH:= _
# replace . with _
ver = $(subst $(DOT),$(DASH),$(VERSION))

.PHONY: build 

build:
	# windows 
	@GOOS=windows GOORCH=386 go build -v -ldflags="-s -w -X 'main.appVersion=$(VERSION)'" -o ddns-client_windows_386_$(ver).exe
	# Linux amd64
	@GOOS=linux GOORCH=amd64 go build -v -ldflags="-s -w -X 'main.appVersion=$(VERSION)'" -o ddns-client_linux_amd64_$(ver)
	# ARMv5 (Raspberry Pi)
	@GOOS=linux GOARCH=arm GOARM=5 go build -v -ldflags="-s -w -X 'main.appVersion=$(VERSION)'" -o ddns-client__linux_amd64_$(ver)
	# MacOS
	@GOOS=darwin GOORCH=386 go build -v -ldflags="-s -w -X 'main.appVersion=$(VERSION)'" -o ddns-client_darwin_386_$(ver)

%:
	@:
