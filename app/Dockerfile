FROM golang:latest AS builder

WORKDIR /app
COPY . /app
RUN go build .

FROM debian:latest AS runner
WORKDIR /app
COPY --from=builder /app/youbine /app/youbine

ENTRYPOINT ["/app/youbine"]