start:
	go run ./cmd/main.go

build:
	go build -o ./bin/messanger/message ./cmd/main.go

docker-build:
	docker build -t message .

lint:
	golangci-lint run ./...

generate:
	go generate ./...

test:
	go test -v ./...