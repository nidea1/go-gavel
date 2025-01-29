FROM golang:1.23.5-alpine AS base

# Install required packages
RUN apk add --no-cache \
    git \
    make \
    protoc \
    protobuf \
    protobuf-dev \
    gcc \
    musl-dev
    
LABEL maintainer="crlidoruk@gmail.com" \
    version="1.0" \
    description="Base image for building"