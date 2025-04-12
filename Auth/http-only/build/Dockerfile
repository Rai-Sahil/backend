FROM golang:1.24.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["./main"]
