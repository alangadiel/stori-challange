FROM golang:1.21
WORKDIR /app
COPY . .

# Download Go modules
RUN go mod download

RUN go run cmd/server.go