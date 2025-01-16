FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy code into the container
COPY ../code/ .
# Download Go modules and dependencies
RUN go mod init local/postgres_app && \
    go mod tidy

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
