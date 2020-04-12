FROM golang:1.14-alpine as builder
RUN apk update && apk add --no-cache git openssh-client

WORKDIR /src
COPY ./ /src
RUN go mod download && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags="-w -s" -v -o /app ./cmd/deployer/main.go 

ENTRYPOINT ["/app"]