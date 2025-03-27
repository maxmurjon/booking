# Builder stage
FROM golang:1.22-alpine AS builder

# Ilova uchun ishchi katalog yaratish
RUN mkdir /app

# Zaruriy paketlarni o'rnatish (FFmpeg uchun kutubxonalar)
RUN apk add --no-cache ffmpeg

# Hamma fayllarni nusxalash
COPY . /app

# Ishchi katalogni sozlash
WORKDIR /app

# Ilovani qurish
RUN go build -o main cmd/main.go

# Minimal tasvir (alpine) yaratish
FROM alpine:3.16

WORKDIR /app

# Zaruriy paketlarni o'rnatish (FFmpeg)
RUN apk add --no-cache ffmpeg

# Qurilgan ilovani nusxalash
COPY --from=builder /app .

# Konfiguratsiya fayli o'zgaruvchisi
ENV DOT_ENV_PATH=config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
