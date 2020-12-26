FROM golang:1-buster AS builder

ARG VERSION

COPY . /workdir

WORKDIR /workdir

RUN go env

RUN go build -v -ldflags="-s -w -X 'main.appVersion=$VERSION'" -o ddns-client

FROM debian:buster-slim

COPY --from=builder /workdir/ddns-client /ddns-client

RUN chmod +x /ddns-client

RUN apt update && apt install ca-certificates -y

ENTRYPOINT [ "/ddns-client" ]
