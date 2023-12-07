FROM golang:latest

WORKDIR /message

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/message ./cmd/main.go

EXPOSE 8080

CMD ["./bin/message"]