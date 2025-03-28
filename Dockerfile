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
COPY api/docs ./docs  # Swagger hujjatlari (endi builder'dan emas, lokal nusxadan olinadi)

# Muhit o‘zgaruvchisini sozlash
ENV DOT_ENV_PATH=config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
