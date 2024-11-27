# syntax=docker/dockerfile:1

# Base image with Go 1.23
FROM golang:1.23-alpine

# Install necessary packages
RUN apk update && apk add --no-cache git

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app/services/notification

# Copy go.mod and go.sum from the project root
COPY go.mod go.sum /app/

# Download dependencies
WORKDIR /app
RUN go mod download

# Copy the entire project into the container
COPY . /app

# Return to the notification service directory
WORKDIR /app/services/notification

# Expose the application port (update if your service uses a different port)
EXPOSE 8080

# Command to run the application with Air
CMD ["air", "-c", ".air.toml"]
