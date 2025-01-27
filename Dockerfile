FROM golang:1.23

COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN go build -o docker_demo

EXPOSE 8080

CMD ["./docker_demo"]

