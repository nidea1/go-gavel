# ===================================
# Build the auth service
# ===================================

FROM go-gavel-base:latest AS auth-builder

# Set the working directory
WORKDIR /app

# Copy the go module files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Download JWT
RUN go get -u github.com/golang-jwt/jwt/v5@v5.2.1

# Copy required files
COPY ./proto/auth/auth.proto ./proto/auth
COPY services/auth services/auth
COPY pkg pkg
COPY scripts scripts

# Generate the protobuf files
RUN chmod +x scripts/generate_protos.sh && ./scripts/generate_protos.sh

# Build the auth service
WORKDIR /app/services/auth
RUN go build -o /auth-service ./cmd/main.go

# ===================================
# Create the final image
# ===================================
FROM alpine:3.21.2

# Set the working directory
WORKDIR /app

# Copy the auth service binary
COPY --from=auth-builder /auth-service .

# Run the auth service
CMD ["./auth-service"]