# Base image
FROM golang:1.21

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first for caching dependencies
COPY go.mod go.sum ./

#Download dependencies
RUN go mod download

# Copy the entire application code
COPY . .

# Build the go application
RUN go build -o main .

# Create a non root user
RUN groupadd -r appgroup && useradd -r -g appgroup appuser

# Change to the non-root user
USER appuser

# Expose the port for external world to use your app
EXPOSE 8080

# Command to run the application
# Run the compiled binary
CMD ["./main"] 