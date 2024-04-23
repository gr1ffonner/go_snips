# Start with a minimal base image containing Go runtime
FROM golang:1.22.0 AS builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the entire project source code into the container
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go
RUN mkdir -p /app/logs

# Start a new stage from scratch
FROM alpine:latest

# Set necessary environment variables for PostgreSQL (if required)
ENV API_KEY=api_key

# Set the current working directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/logs ./logs

CMD ["./main"] >> /app/logs/app.log 2>&1
