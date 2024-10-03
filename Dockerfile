# Step 1: Build Stage
FROM golang:1.22-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker's caching mechanism for dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application statically to avoid runtime dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/my-go-app

# Step 2: Minimal Final Image
FROM scratch

# Copy the statically compiled binary from the build stage
COPY --from=builder /app/bin/my-go-app /app/my-go-app

# Command to run the binary
ENTRYPOINT ["/app/my-go-app"]
