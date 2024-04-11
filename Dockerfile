# Start from the official Go image to build our executable
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

# Use `go build` to compile the binary executable of our Go program
RUN go mod init myserver && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file and .env file from the previous stage
COPY --from=builder /app/server .
COPY .env .

# Copy your Makefile or any other necessary files to the image
COPY Makefile .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
