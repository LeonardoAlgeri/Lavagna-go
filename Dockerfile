# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -mod vendor -o main .

FROM debian:buster-slim
COPY --from=0 /app/main /app/
WORKDIR /app

# Add Maintainer Info
LABEL maintainer="Leonardo Algeri"

RUN useradd -u 8877 lavagna-go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

# Change to non-root privilege
USER lavagna-go
