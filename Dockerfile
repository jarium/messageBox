# Build stage
FROM golang:latest AS builder
WORKDIR /app

# Environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o messagebox ./cmd/main.go

# Build the custom health check binary
RUN go build -o healthcheck ./cmd/healthcheck.go

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/messagebox .
COPY --from=builder /app/healthcheck .

# Run the binary
CMD ["./messagebox"]
