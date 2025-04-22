FROM golang:1.24.0-alpine3.21 AS builder
    WORKDIR /app

    COPY ./orderService/ ./

    RUN go build -o orderService .

FROM alpine:3.21.3

    WORKDIR /app
    RUN apk add make
    COPY --from=builder /app/orderService .
    COPY --from=builder /app/migrations ./migrations
    COPY --from=builder /app/Makefile ./Makefile
    COPY ./bin/goose /usr/local/bin/goose

    RUN chmod +x /usr/local/bin/goose

    EXPOSE 8082

    CMD ["./orderService"]

