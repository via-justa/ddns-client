FROM golang:1-buster AS builder

ENV VERSION=0.1.0

COPY . /workdir

WORKDIR /workdir

RUN go build -v -ldflags="-s -w -X 'main.appVersion=$VERSION'" -o ddns-client

FROM debian:buster-slim

ENV HETZNER_API_KEY=

COPY --from=builder /workdir/ddns-client /ddns-client

RUN chmod +x /ddns-client

RUN apt update && apt install ca-certificates -y

ENTRYPOINT [ "/ddns-client" ]
