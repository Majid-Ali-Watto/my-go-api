# Use the Go 1.23 image as the build stage
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app from the cmd/app directory
RUN go build -o my-go-api ./cmd/app/main.go

# Start a new stage from a lightweight alpine image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/my-go-api .

# Expose port 8080 to the outside world
EXPOSE 8080

# Ensure the binary is executable
RUN chmod +x ./my-go-api

# Command to run the executable
CMD ["./my-go-api"]
