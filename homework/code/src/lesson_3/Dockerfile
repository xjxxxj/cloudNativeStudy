FROM golang:alpine AS builder

WORKDIR /build

COPY httpserver.go .

RUN go build -o httpserver httpserver.go



FROM alpine

WORKDIR /app

ENV VERSION='0.1'

COPY --from=builder /build/httpserver /app/httpserver

CMD ["./httpserver"]
