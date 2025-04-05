FROM golang:1.24.0-alpine3.21 AS builder

    WORKDIR /app

    COPY . .

    RUN go mod init catalogService && go mod tidy

    RUN go build -o catalogService .

FROM alpine:3.21.3

    WORKDIR /app

    COPY --from=builder /app/catalogService .
    COPY --from=builder /app/migrations ./migrations

    EXPOSE 8080

    CMD ["./catalogService"]

