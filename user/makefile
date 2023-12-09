start:
	go run ./cmd/main.go

build:
	go build -o ./bin/messanger/user ./cmd/main.go

docker-build:
	docker build -t user .

lint:
	golangci-lint run ./...

generate:
	go generate ./...

test:
	go test -v ./...