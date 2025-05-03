FROM golang:1.24.2-bookworm AS builder
    WORKDIR /app

    COPY ./backgroundWorkerService/ ./

    RUN go build -o backgroundWorkerService .

FROM debian:bookworm

    WORKDIR /app
    RUN apt update && apt install make
    COPY --from=builder /app/backgroundWorkerService .
    COPY --from=builder /app/migrations ./migrations
    COPY --from=builder /app/Makefile ./Makefile
    COPY ./bin/goose /usr/local/bin/goose

    RUN chmod +x /usr/local/bin/goose

    CMD ["./backgroundWorkerService"]

