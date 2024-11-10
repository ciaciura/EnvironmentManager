# Use the official Golang image as the base
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /

# Copy the Go module files first to optimize build caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app main.go

FROM alpine:latest

COPY --from=builder /app /app

CMD [ "/app" ]
#CMD ["sleep", "infinity"]
