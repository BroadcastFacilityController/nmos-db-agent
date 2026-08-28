# Specifies a parent image
FROM golang:1.26.6

WORKDIR /app

# Copy Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./
COPY nmosdb/ ./nmosdb/
COPY watchdog/ ./watchdog/

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -o /bfc-nmos-db-agent-2

# Inbound Ports
# EXPOSE 8080

# Start
CMD [ "/bfc-nmos-db-agent-2" ]