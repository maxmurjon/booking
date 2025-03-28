# 1. Builder stage
FROM golang:1.22-alpine AS builder  # 1.23 emas
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Swagger generatsiya qilish
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g cmd/main.go -o docs

# Statik binary qurish
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o main cmd/main.go

# 2. Final stage
FROM alpine:3.19  # Versiyani belgilang
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs
ENV DOT_ENV_PATH=/app/config/.env
CMD ["/app/main"]