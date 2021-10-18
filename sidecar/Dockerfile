ARG VERSION=v0.1.0

FROM golang:1.17-alpine

WORKDIR /

# copy mod and sum files
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o cmd cmd/main.go

# Exposes port 8081
EXPOSE 8081

# Run the binary
CMD ["./cmd/main"]