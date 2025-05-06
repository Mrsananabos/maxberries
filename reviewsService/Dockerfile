FROM golang:1.24.2-bookworm AS builder
    WORKDIR /app

    COPY ./reviewsService/ ./

    RUN go build -o reviewsService .

FROM debian:bookworm

    WORKDIR /app

    COPY --from=builder /app/reviewsService .

    CMD ["./reviewsService"]

