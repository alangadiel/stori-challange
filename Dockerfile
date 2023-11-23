FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

# Download Go modules
RUN go mod download

# Build the Go app
RUN go build -o /server cmd/server.go

# Run the executable
CMD [ "/server" ]