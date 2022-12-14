FROM golang:latest AS builder

COPY ./app /app
WORKDIR /app
RUN go build .

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/youbine /app/

ENTRYPOINT ["./youbine"]