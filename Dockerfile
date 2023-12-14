# Use an official Golang runtime as a base image
FROM golang:1.21-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a lightweight Alpine image as the final base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/ .

# Expose the port on which the Go application will run
EXPOSE 6969 

# Command to run the application
CMD ["./main"]
