# Build stage
FROM golang:latest AS builder
WORKDIR /app

# Environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o messagebox ./cmd

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/messagebox .

# Run the binary
CMD ["./messagebox"]
