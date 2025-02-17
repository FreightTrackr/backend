# Gunakan image Go sebagai base image
FROM golang:1.22.9-alpine AS builder

# Install dependencies yang diperlukan
RUN apk add --no-cache git

# Set environment variable
ENV GO111MODULE=on

# Buat direktori kerja di dalam container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh kode sumber
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main.exe main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main.exe .

CMD ["./main.exe"]