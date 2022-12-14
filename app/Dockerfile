FROM golang:latest AS builder

WORKDIR /app
COPY . /app
RUN go build .

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/youbine /app/

ENTRYPOINT ["./youbine"]