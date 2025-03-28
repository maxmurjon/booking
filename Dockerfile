# 1. Builder stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Kerakli kutubxonalarni o‘rnatish
RUN apk add --no-cache git ca-certificates

# Go mod fayllarni yuklash
COPY go.mod go.sum ./
RUN go mod download

# Loyihani nusxalash
COPY . .

# Swagger hujjatlarini yaratish (Agar kerak bo‘lsa)
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init -g cmd/main.go -o api/docs

# Go ilovasini build qilish
RUN go build -o main cmd/main.go

# 2. Final stage
FROM alpine:3.16

WORKDIR /app

# Kerakli kutubxonalarni o‘rnatish
RUN apk add --no-cache ca-certificates

# Qurilgan Go ilovasini nusxalash
COPY --from=builder /app/main .  # Asosiy Go dastur
COPY --from=builder /app/config ./config  # Konfiguratsiya fayllari
COPY --from=builder /app/api/docs ./docs  # Swagger hujjatlari

# Muhit o‘zgaruvchisini sozlash
ENV DOT_ENV_PATH=config/.env

CMD ["/app/main"]
