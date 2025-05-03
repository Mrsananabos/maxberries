FROM golang:1.24.0-alpine3.21 AS builder
    WORKDIR /app

    COPY ./authService/ ./

    RUN go build -o authService .

FROM alpine:3.21.3

    WORKDIR /app

    COPY --from=builder /app/authService .

    CMD ["./authService"]

