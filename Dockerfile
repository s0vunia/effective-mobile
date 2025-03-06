ARG CONFIG_PATH

FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./bin/crud_server cmd/worker/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/crud_server .
COPY .env .

CMD ["./crud_server"]
