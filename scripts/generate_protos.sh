#!/usr/bin/env bash

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

# Project proto directory
BASE_PROTO_DIR="proto"

# Loop over all subdirectories under proto/
for dir in "$BASE_PROTO_DIR"/*; do
    # Check if it's a directory
    if [ -d "$dir" ]; then
        echo "Processing directory: $dir"
        
        # List .proto files in that directory (remove -maxdepth if you need to check nested subdirectories)
        proto_files=$(find "$dir" -maxdepth 1 -name "*.proto")
        
        # If there are .proto files, compile them
        if [ -n "$proto_files" ]; then
            echo "Files to compile: $proto_files"
            protoc \
              -I="$BASE_PROTO_DIR" \
              -I="$dir" \
              --go_out="$dir" \
              --go_opt=paths=source_relative \
              --go-grpc_out="$dir" \
              --go-grpc_opt=paths=source_relative \
              $proto_files
            echo "Compilation completed: $dir"
        else
            echo "No .proto file found in $dir."
        fi
    fi
done

echo "All proto files compiled successfully."
