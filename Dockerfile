# Use an official Golang base image
FROM golang:1.21.0-alpine3.18

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules files to the container
# COPY go.mod go.sum ./
COPY go.mod ./

# Download and cache Go modules dependencies
RUN go mod download

# Copy the rest of the application files to the container
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port your Golang API will listen on (change if needed)
EXPOSE 8050

# Command to run the Golang API application
CMD ["./app"]
