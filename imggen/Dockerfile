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
COPY bg/ ./bg/ 
COPY PtSerif_Regular.ttf ./ 

# Download and install Go dependencies
RUN go mod download

# Copy the entire project source code into the container
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set necessary environment variables for PostgreSQL (if required)
ENV DB_HOST=host
ENV API_KEY=api_key
ENV API_SECRET_KRAKEN=api_secret_kraken
ENV API_KEY_KRAKEN=api_key_kraken

# Set the current working directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/bg /app/bg/
COPY --from=builder /app/PtSerif_Regular.ttf /app/ 

# Expose port 8080 to the outside world
EXPOSE 8080
CMD ["./main"]
