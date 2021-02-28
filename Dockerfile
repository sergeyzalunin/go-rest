FROM golang:1.15.8-alpine

WORKDIR /opt/code
ADD . /opt/code

RUN apk update && apk upgrade \
    && apk add --no-cache git \
    && apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz \
    && ./migrate.linux-amd64 -path migrations -database $DATABASE_URL -verbose up

RUN go mod download

RUN go build -o bin/go-rest ./cmd/server

ENTRYPOINT [ "bin/go-rest" ]