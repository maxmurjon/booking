# Builder stage
FROM golang:1.22-alpine AS builder

# Ilova uchun ishchi katalog yaratish
RUN mkdir /app

# Hamma fayllarni nusxalash
COPY . /app

# Ishchi katalogni sozlash
WORKDIR /app

# Ilovani qurish
RUN go build -o main cmd/main.go

# Minimal tasvir (alpine) yaratish
FROM alpine:3.16

WORKDIR /app

# Qurilgan ilovani nusxalash
COPY --from=builder /app/main .  # Faqat kerakli fayllarni olib kelish
COPY --from=builder /app/api/docs ./api/docs  # Swagger fayllarini qo‘shish

# Konfiguratsiya fayli o'zgaruvchisi
ENV DOT_ENV_PATH=config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
