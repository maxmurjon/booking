# 1. Builder stage
FROM golang:1.22-alpine AS builder  # 1.23 emas, 1.22 yoki 1.21

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Swagger docs manzilini to'g'rilash
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g cmd/main.go -o docs  # api/docs emas, direktoriyani soddalashtirish

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go  # CGO ni o'chirish

# 2. Final stage
FROM alpine:3.19  # Yangi versiya

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs  # To'g'ri manzil

ENV DOT_ENV_PATH=config/.env

CMD ["/app/main"]