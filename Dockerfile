# Builder stage
FROM golang:1.23.0-alpine AS builder

# Ishchi katalogni sozlash
WORKDIR /app

# Muhit o‘zgaruvchilarini sozlash
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Faqat kerakli fayllarni nusxalash
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ilovani qurish
RUN go build -o main cmd/main.go

# Minimal tasvir yaratish
FROM alpine:latest

WORKDIR /app

# Kerakli fayllarni nusxalash
COPY --from=builder /app/main .
COPY --from=builder /app/config/.env config/.env

# Ilovani ishga tushirish
CMD ["/app/main"]
