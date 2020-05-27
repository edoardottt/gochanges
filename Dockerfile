# Start from the latest golang base image
FROM golang:latest

# Add env vars
ENV FILE_NAME="example/example1.txt"
ENV MONGO_CONN="mongodb://172.17.0.1:27017"
ENV DB_NAME="gochangesdb"


# Add Maintainer Info
LABEL maintainer="edoardottt <edoardott@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /gochanges

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Run the Go app
CMD ["go","run","."]