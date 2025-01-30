#!/usr/bin/env bash

# Root directory containing source .proto files
PROTO_SRC_DIR="proto"

# Directory where generated .pb.go files will be placed
PROTO_OUT_DIR="proto-gen"

# Create proto-gen directory (doesn't error if exists, creates if not)
mkdir -p "$PROTO_OUT_DIR"

# Find all .proto files
PROTO_FILES=$(find "$PROTO_SRC_DIR" -type f -name "*.proto")

# Loop through each .proto file
for proto_file in $PROTO_FILES; do
    # Example: proto_file = "proto/auth/auth.proto"
    # subdir = "/auth"
    subdir=$(dirname "$proto_file" | sed "s|^$PROTO_SRC_DIR||")

    # Create the same subdirectories in output directory
    # Example: "proto-gen/auth"
    mkdir -p "$PROTO_OUT_DIR/$subdir"

    echo "Generating $proto_file into $PROTO_OUT_DIR/$subdir"

    # protoc command
    protoc \
        -I="$PROTO_SRC_DIR" \
        --go_out="$PROTO_OUT_DIR" \
        --go_opt=paths=source_relative \
        --go-grpc_out="$PROTO_OUT_DIR" \
        --go-grpc_opt=paths=source_relative \
        "$proto_file"
done

echo "All protos have been successfully generated!"
