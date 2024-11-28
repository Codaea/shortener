FROM golang:1.23.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o api ./cmd

# Use Alpine as the base image for the final stage
FROM alpine:3.18

# Install necessary dependencies
RUN apk add --no-cache libc6-compat

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api /app/api

# Command to run the executable
CMD ["/app/api"]