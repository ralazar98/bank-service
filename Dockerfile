FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .. .

WORKDIR /app/cmd

RUN go build -o bank-service .

EXPOSE 8080

CMD ["./bank-service"]