FROM golang:1.24.0-alpine3.21 AS builder
    WORKDIR /app

    COPY ./reviewsService/ ./

    RUN go build -o reviewsService .

FROM alpine:3.21.3

    WORKDIR /app

    COPY --from=builder /app/reviewsService .

    CMD ["./reviewsService"]

