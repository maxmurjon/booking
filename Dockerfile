# 1. Builder stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Kerakli kutubxonalarni o‘rnatish
RUN apk add --no-cache git ca-certificates

# Go mod fayllarni yuklash
COPY go.mod go.sum ./
RUN go mod download

# Swagger CLI ni o‘rnatish
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2

# Loyihani nusxalash
COPY . .

# Swagger hujjatlarini generatsiya qilish
RUN mkdir -p /app/docs
RUN swag init -g cmd/main.go -o /app/docs

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
COPY --from=builder /app/docs ./docs  # Swagger hujjatlari

# Muhit o‘zgaruvchisini sozlash
ENV DOT_ENV_PATH=config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
