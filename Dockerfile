FROM golang:1.15.8-alpine

WORKDIR /opt/code
ADD . /opt/code

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/go-rest ./cmd/server

ENTRYPOINT [ "bin/go-rest" ]