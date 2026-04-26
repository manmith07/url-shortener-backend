# Use Go image
FROM golang:1.26

# Set working directory
WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build app
RUN go build -o main ./cmd

# Expose port
EXPOSE 8080

# Run app
CMD ["./main"]
