FROM golang:1.23

# Copy the go.mod and go.sum files to the /build directory
COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN go build -o docker_demo

# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["./docker_demo"]

