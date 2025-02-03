# ===================================
# Build the auth service
# ===================================

FROM go-gavel-base:latest AS auth-builder

# Set the working directory
WORKDIR /app

# Copy required files
COPY proto/go.mod ./proto/
COPY proto/auth   ./proto/auth/
COPY proto/common ./proto/common/
COPY services/auth  ./services/auth/
COPY pkg            ./pkg/
COPY scripts        ./scripts/

# Generate the protobuf files
RUN chmod +x scripts/generate_protos.sh && ./scripts/generate_protos.sh

# Workspace setup
RUN go work init \
    ./services/auth \
    ./pkg \
    ./proto \
    && go work sync

# Build the auth service
WORKDIR /app/services/auth
RUN go mod tidy && \
    go mod download && \
    go build -o auth-service ./cmd

# ===================================
# Create the final image
# ===================================
FROM alpine:3.21.2

# Set the working directory
WORKDIR /app

# Copy the auth service binary
COPY --from=auth-builder /app/services/auth/auth-service .

# Expose the port
EXPOSE 8080

# Run the auth service
CMD ["./auth-service"]