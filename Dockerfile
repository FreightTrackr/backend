# Gunakan image Go sebagai base image
FROM golang:1.20-alpine

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
EXPOSE 8080

ENTRYPOINT ["/app"]
CMD ["./main"]