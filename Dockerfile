# Stage 1: Build the Go app
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy only go.mod and go.sum to cache dependency download step
COPY go.mod go.sum ./

# Download Go dependencies, caching the result
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create a minimal image to run the application
FROM alpine:3.18

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port that the app listens on
EXPOSE 8005

# Command to run the app
CMD ["./main"]
