# Gunakan image Go sebagai base image
FROM golang:1.21.3-alpine

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
RUN go build -o main .

# Ekspose port untuk aplikasi
EXPOSE 3000

# Tentukan entrypoint dan command
ENTRYPOINT ["/app/main"]

# Menggunakan port dari environment variable yang diatur oleh Cloud Run
CMD ["-port", "3000"]