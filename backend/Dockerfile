# Use official Golang image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy Go module files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the backend code
COPY . .

# Build binary
RUN go build -o server .

# Expose port and run
EXPOSE 8000
CMD ["./server"]
