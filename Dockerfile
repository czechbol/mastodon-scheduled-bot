# Start from the latest golang base image
FROM golang:1.21-alpine as builder

RUN apk --no-cache add ca-certificates tzdata && update-ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mastodon-siren .

######## Start a new stage from scratch #######
FROM scratch  

# Copy the tzdata and ca-certificates from the builder image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/mastodon-siren .

# Command to run the executable
ENTRYPOINT ["./mastodon-siren"]