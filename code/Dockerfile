FROM golang:1.24-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY src/go.mod src/go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY src .

# Build the Go application
RUN go build -o forum ./cmd

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./forum"]

