# 1. Builder stage
FROM golang:1.23-alpine AS builder

# Muhitni sozlash
WORKDIR /app

# Go va kerakli paketlarni o‘rnatish
RUN apk add --no-cache git

# Go mod fayllarni nusxalash va bog‘liqliklarni yuklash
COPY go.mod go.sum ./
RUN go mod tidy

# Swagger CLI ni o‘rnatish
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Hamma loyihani nusxalash
COPY . .

# Swagger hujjatlarini generatsiya qilish
RUN swag init -g cmd/main.go

# Ilovani build qilish
RUN go build -o main cmd/main.go

# 2. Final stage
FROM alpine:3.16

WORKDIR /app

# Qurilgan Go ilovasini nusxalash
COPY --from=builder /app/main .  # Asosiy Go dastur
COPY --from=builder /app/config .  # Konfiguratsiya fayllari
COPY --from=builder /app/api/docs ./api/docs  # Swagger hujjatlari

# `.env` fayl yo‘qligi sababli konfiguratsiyani muhit o‘zgaruvchilari orqali sozlash
ENV DOT_ENV_PATH=config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
