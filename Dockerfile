# Build stage with Go 1.24
FROM golang:1.24 as builder

# Set environment to build a static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire app
COPY . .

# Build static binary
RUN go build -o main .

# Final image (scratch = smallest possible container)
FROM scratch

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/main .

EXPOSE 8080

# Run the binary
CMD ["./main"]
