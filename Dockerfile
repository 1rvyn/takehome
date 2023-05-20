# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Env vars for hashing / other needed things
ENV SECRET_KEY=$SECRET_KEY \
    SALT=$SALT \
    SECRET_KEY2=$SECRET_KEY2

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]