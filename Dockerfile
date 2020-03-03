# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Leonardo Algeri"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Form directory
ADD form ./form

# Copy go mod and sum files
COPY main.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get -u github.com/go-sql-driver/mysql

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
