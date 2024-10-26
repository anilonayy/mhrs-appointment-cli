FROM golang:1.23-alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o mhrs-cli .

FROM alpine:latest

COPY --from=builder /app/mhrs-cli /usr/local/bin/mhrs-cli
COPY .env .env

CMD ["mhrs-cli"]