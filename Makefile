build:
	go build -o go-rest -v ./cmd/server

run: build
	./go-rest	

test:
	go test -v -race -timeout 30s ./...

_DEFAULT_GO := run