# syntax=docker/dockerfile:1

FROM golang:1.26

# Set work directory / destination for copy
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./
COPY watchdog/  ./watchdog/
COPY database/ ./database/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /nmos-db-agent

# Run
CMD ["/nmos-db-agent"]