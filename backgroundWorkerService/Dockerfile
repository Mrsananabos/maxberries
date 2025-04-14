FROM golang:1.24.0-alpine3.21 AS builder
    WORKDIR /app

    COPY ./backgroundWorkerService/ ./

    RUN go build -o backgroundWorkerService .

FROM alpine:3.21.3

    WORKDIR /app

    COPY --from=builder /app/backgroundWorkerService .

    CMD ["./backgroundWorkerService"]

