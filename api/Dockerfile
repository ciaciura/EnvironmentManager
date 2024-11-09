# Use the official Golang image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app/api

# Copy the Go module files first to optimize build caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o goapp main.go

# Expose the application's port
EXPOSE 8000

# Run the Go application
CMD ["./goapp"]
#CMD ["sleep", "infinity"]
