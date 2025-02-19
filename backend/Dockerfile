# Use the official Golang image
FROM golang:1.23-alpine AS backend

WORKDIR /app

# Copy only go.mod and go.sum first (for better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Run the app
CMD ["./main"]


