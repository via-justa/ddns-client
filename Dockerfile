FROM golang:1-alpine AS builder

ARG VERSION

COPY . /workdir

WORKDIR /workdir

RUN go env
RUN go build -v -ldflags="-s -w -X 'main.appVersion=$VERSION'" -o ddns-client

FROM alpine

COPY --from=builder /workdir/ddns-client /ddns-client

RUN chmod +x /ddns-client

RUN apk add ca-certificates

ENTRYPOINT [ "/ddns-client" ]
