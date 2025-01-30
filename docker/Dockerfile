FROM golang:1.20
ARG TIMEZONE="Asia/Dhaka"
# Timezone
ENV TZ="${TIMEZONE}"
RUN apt-get update && apt-get install -y tzdata && \
    ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy code into the container
COPY code/ .
#COPY . .
# Download Go modules and dependencies
RUN go mod init local/postgres_app && \
    go mod tidy

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
