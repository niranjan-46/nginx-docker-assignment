# -------------------------------
# Build stage
# -------------------------------
FROM golang:1.21-alpine AS builder

# Install git (needed for go modules with external repos)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod (go.sum is optional, but we skip it if missing)
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# -------------------------------
# Final stage
# -------------------------------
FROM alpine:latest

# Add curl for healthcheck
RUN apk --no-cache add curl

# Set working directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port 8001
EXPOSE 8001

# Healthcheck for Docker to know if the container is alive
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD curl -f http://localhost:8001/health || exit 1

# Start the service
CMD ["./main"]

