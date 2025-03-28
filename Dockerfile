# === Builder stage ===
FROM golang:1.22-alpine AS builder

# Ishchi katalogni yaratish
WORKDIR /app

# Go mod fayllarni nusxalash va bog‘liqliklarni yuklash
COPY go.mod go.sum ./
RUN go mod tidy

# Hamma fayllarni nusxalash
COPY . .

# `swag` ni o‘rnatish va Swagger hujjatlarini yaratish
RUN go install github.com/swaggo/swag/cmd/swag@latest \
    && /go/bin/swag init -g cmd/main.go

# Ilovani qurish
RUN go build -o main cmd/main.go

# === Minimal image (alpine) ===
FROM alpine:3.16

WORKDIR /app

# Qurilgan ilovani va kerakli fayllarni nusxalash
COPY --from=builder /app/main .  # Asosiy Go dastur
COPY --from=builder /app/config .  # Konfiguratsiya fayllari
COPY --from=builder /app/api/docs ./api/docs  # Swagger hujjatlari

# `.env` fayl yo‘qligi sababli konfiguratsiyani muhit o‘zgaruvchilari orqali sozlash
ENV DOT_ENV_PATH=config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
