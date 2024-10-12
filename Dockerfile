FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bank-service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bank-service .
EXPOSE 8080
CMD ["./bank-service"]