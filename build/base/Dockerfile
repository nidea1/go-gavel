FROM golang:1.23.5-alpine AS builder-base

# Install required packages
RUN apk add --no-cache \
    add \
    gcc \
    g++ \
    make \
    ca-certificates \
    protoc \
    protobuf \
    protobuf-dev

LABEL maintainer="crlidoruk@gmail.com" \
    version="1.0" \
    description="Base image for building"