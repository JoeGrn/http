FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o ./http cmd/http/main.go

EXPOSE 8080

CMD ["./http"]