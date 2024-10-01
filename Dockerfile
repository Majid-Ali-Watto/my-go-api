# Use a base image with the desired Go version
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only the necessary files to build the app
COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/app/ ./cmd/app/

# Build the Go app
RUN go build -o my-go-api ./cmd/app/main.go

# Start a new stage from scratch using a lightweight image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/my-go-api .

# Command to run the executable
CMD ["./my-go-api"]
