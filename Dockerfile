# 1. Builder stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g cmd/main.go -o docs
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# 2. Final stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs
ENV DOT_ENV_PATH=/app/config/.env
CMD ["/app/main"]