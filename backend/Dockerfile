# Use Golang as the base image
FROM golang:1.18

# Set the working directory
WORKDIR /usr/src/app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o server .

# Expose port 8080
EXPOSE 8080

# Start the WebSocket server
CMD ["./server"]
