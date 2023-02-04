FROM golang:latest AS builder

WORKDIR /app
COPY . /app
RUN go build .

FROM debian:latest AS runner
WORKDIR /app
COPY --from=builder /app/youbine /app/youbine
RUN apt-get update && apt-get install -y \
    ca-certificates \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/app/youbine"]