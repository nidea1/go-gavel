#!/usr/bin/env bash

# Stop the script on error
set -e

# Check if protoc is installed
if ! command -v protoc &> /dev/null
then
    echo "protoc not found. Please install protoc first."
    exit 1
fi

# Check if protoc-gen-go is installed
if ! command -v protoc-gen-go &> /dev/null
then
    echo "protoc-gen-go not found."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    echo "protoc-gen-go installed."
fi

# Check if protoc-gen-go-grpc is installed
if ! command -v protoc-gen-go-grpc &> /dev/null
then
    echo "protoc-gen-go-grpc not found."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    echo "protoc-gen-go-grpc installed."
fi

# Use fixed path for Docker environment
PROJECT_ROOT="/app"
echo "PROJECT_ROOT: $PROJECT_ROOT"
BASE_PROTO_DIR="$PROJECT_ROOT/proto"

# Loop over all subdirectories under proto/
find "$BASE_PROTO_DIR" -type d | while read dir; do
    echo "Checking directory: $dir"
    
    # Find all proto files in the current directory
    proto_files=$(find "$dir" -maxdepth 1 -name "*.proto")
    
    if [ -n "$proto_files" ]; then
        echo "Compiling proto files in: $dir"
        echo "Proto files to compile: $proto_files"
        
        # Get relative path from BASE_PROTO_DIR to current directory
        rel_dir=${dir#"$BASE_PROTO_DIR/"}
        
        # Generate Go code for each proto file
        cd "$BASE_PROTO_DIR" && protoc \
            --proto_path="." \
            --go_out="." \
            --go_opt=paths=source_relative \
            --go-grpc_out="." \
            --go-grpc_opt=paths=source_relative \
            $(find "$rel_dir" -maxdepth 1 -name "*.proto")

        if [ $? -eq 0 ]; then
            echo "✓ Successfully compiled protos in: $dir"
            echo "Generated files in directory:"
            ls -la "$dir"/*.pb.go || true
        else
            echo "✗ Failed to compile protos in: $dir"
            exit 1
        fi
    else
        echo "No .proto files found in: $dir"
    fi
done

echo "✓ All proto files compiled successfully!"
