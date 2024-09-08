# Stage 1: Build the Go binary
FROM golang:1.20-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the index.html file
COPY index.html ./

# Expose the port the application runs on
EXPOSE 8005

# Command to run the Go application
CMD ["./main"]
