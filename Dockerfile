# Step 1: Build the Go binary
FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o ./http cmd/http/main.go

EXPOSE 8080

CMD ["./http"]