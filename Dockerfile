# 🛠️ --- Stage 1: Build ---
    FROM golang:1.20 AS builder

    WORKDIR /app
    
    # Copy go.mod and download dependencies
    COPY go.mod ./
    RUN go mod tidy
    
    # Copy source code
    COPY . .
    
    # Ensure Go compiles a static binary (important for Alpine)
    RUN CGO_ENABLED=0 go build -o /app/blueis .
    
    # 🐋 --- Stage 2: Minimal runtime ---
    FROM alpine:latest
    
    WORKDIR /app
    
    # Copy the compiled binary from the builder stage
    COPY --from=builder /app/blueis /app/blueis
    
    # Ensure the binary is executable
    RUN chmod +x /app/blueis
    
    # Expose the required port
    EXPOSE 7171
    
    # Run the application
    CMD ["/app/blueis"]
    