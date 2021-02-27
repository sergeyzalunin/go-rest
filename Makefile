db?="postgresql://admin:123@172.28.1.5/restapi_dev?sslmode=disable"
name?=empty
id?=0

build:
	go build -o go-rest -v ./cmd/server

run: build
	./go-rest	

test:
	go test -race -timeout 30s ./...

testv:
	go test -v -race -timeout 30s ./...

killshim:
	sudo killall containerd-shim

dc-up:
	docker-compose up --remove-orphans

dc-down: killshim
	docker-compose down

migr-up:
	migrate -path migrations -database ${db} -verbose up

migr-down:
	migrate -path migrations -database ${db} -verbose down

migr-create:
	migrate create -ext sql -dir migrations -seq ${name}

migr-rollback:
	migrate -path migrations -database ${db} -verbose force ${id}

pgexec:
	docker exec -it postgres psql -U admin -W 123 -d restapi_dev -h 172.28.1.5 -p 5432

createuser:
	curl -X 'POST' -H 'Content-Type: application/json' --data '{"email":"test1@ree.ru","password":"123456"}'  http://localhost:8080/users

_DEFAULT_GO := run