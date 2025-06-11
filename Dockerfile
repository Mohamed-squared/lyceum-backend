# Stage 1: Build the Go binary
# UPDATED THIS LINE from 1.22 to 1.24
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application.
# -o /app/server specifies the output path for the binary.
# CGO_ENABLED=0 disables Cgo, creating a static binary.
# -ldflags="-w -s" strips debugging information, reducing binary size.
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /app/server ./cmd/server

# Stage 2: Create the final, minimal image
FROM alpine:latest

# It's good practice to run as a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /home/appuser/

# Copy the built binary from the 'builder' stage
COPY --from=builder /app/server .

# This command will be run when the container starts.
# It executes our compiled Go application.
CMD ["./server"]
