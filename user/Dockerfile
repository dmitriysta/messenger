FROM golang:latest

WORKDIR /user

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/user ./cmd/main.go

EXPOSE 8081

CMD ["./bin/user"]