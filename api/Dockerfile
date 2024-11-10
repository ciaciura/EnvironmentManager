# Use the official Golang image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first to optimize build caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Run the Go application
#CMD ["/api"]
CMD ["sleep", "infinity"]
