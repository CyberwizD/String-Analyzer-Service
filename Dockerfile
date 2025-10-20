# Use the official Golang image as the base image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o server cmd/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./cmd/main"]