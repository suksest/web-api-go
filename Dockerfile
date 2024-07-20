FROM golang:1.22.3-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o webapp

FROM alpine:3.20 AS release
WORKDIR /app
COPY --from=builder /app/webapp .
EXPOSE 8000

ENTRYPOINT ["./webapp"]