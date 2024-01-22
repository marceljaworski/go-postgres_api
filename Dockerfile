# A microservice in Go packaged into a container image.
FROM golang

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code.
COPY *.go ./

# Build
RUN go build -o /go-postgres_api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
EXPOSE 8081

CMD ["/go-postgres_api"]