# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary dependencies for building the Go application
RUN apk add --no-cache git gcc musl-dev

# Copy the backend directory contents
COPY backend/ .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o quizex ./cmd

# Stage 2: Create the final image using Alpine
FROM alpine:latest

# Install SQLite dependencies for Alpine
RUN apk add --no-cache sqlite sqlite-dev

# Create app directory
WORKDIR /app

# Copy necessary files and directories from builder
COPY --from=builder /app/cmd/view ./cmd/view
COPY --from=builder /app/uploads ./uploads
COPY --from=builder /app/quizex . 
COPY --from=builder /app/data/migrations ./migrations 
COPY --from=builder /app/database ./database
COPY backend/.env . 

# Set volume for SQLite database
VOLUME ["/app/database"]

# Expose the application port
EXPOSE 9090

# Run the application
CMD ["./quizex"]
