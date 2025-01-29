FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application (static binary)
RUN go build -o task-manager-app ./cmd/main.go

# Final Stage: Create a minimal production image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/task-manager-app .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the compiled application
CMD ["./task-manager-app"]