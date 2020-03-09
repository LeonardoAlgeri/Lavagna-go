# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Leonardo Algeri"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY main.go ./

# Download all dependencies & set user
RUN go get -u github.com/go-sql-driver/mysql github.com/gorilla/mux

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

FROM debian:buster-slim
COPY --from=0 /app /app/
RUN useradd -u 8877 lavagna-go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

# Change to non-root privilege
USER lavagna-go
