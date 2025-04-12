FROM golang:1.24.0-alpine3.21 AS builder
    WORKDIR /app

    COPY ./catalogService/ ./

    RUN go build -o catalogService .

FROM alpine:3.21.3

    WORKDIR /app
    RUN apk add make
    COPY --from=builder /app/catalogService .
    COPY --from=builder /app/migrations ./migrations
    COPY --from=builder /app/Makefile ./Makefile
    COPY ./bin/goose /usr/local/bin/goose

    RUN chmod +x /usr/local/bin/goose

    EXPOSE 8080

    CMD ["./catalogService"]

