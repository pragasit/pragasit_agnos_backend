# FROM golang:latest

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o main

# CMD ["./main"]

# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the local package files to the container's workspace
COPY . .

# Download and install any needed dependencies


# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go application
CMD ["./main"]
