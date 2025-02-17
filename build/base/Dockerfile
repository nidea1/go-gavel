FROM golang:1.24-alpine AS builder-base

# Install required packages
RUN apk add --no-cache \
    gcc \
    g++ \
    make \
    bash \
    ca-certificates \
    protoc \
    protobuf \
    protobuf-dev

LABEL maintainer="crlidoruk@gmail.com" \
    version="1.0" \
    description="Base image for building"