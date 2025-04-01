# 🛠️ --- Stage 1: Build ---
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Copy source code separately
COPY . .

# Build the Go application
RUN go build -o blueis .

# 🐋 --- Stage 2: Minimal runtime ---
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/blueis .

# Expose the port the app listens on
EXPOSE 7171

# Command to run the app
CMD ["./main.go"]
