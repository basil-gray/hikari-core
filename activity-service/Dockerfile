FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main ./cmd

FROM alpine:latest
COPY --from=builder /app/main /app/main
CMD ["/app/main"]