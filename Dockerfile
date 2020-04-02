FROM golang:1.14-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /src
COPY ./ /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags="-w -s" -v -o /src/app ./cmd/deployer/main.go 

ENTRYPOINT ["/src/app"]