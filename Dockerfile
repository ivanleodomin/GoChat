FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/api/main.go

CMD ["/app/main"]