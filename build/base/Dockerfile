FROM golang:1.23.5-alpine AS builder-base

# Install required packages
RUN apk add --no-cache \
    git \
    bash \
    make \
    protoc \
    protobuf \
    protobuf-dev \
    gcc \
    build-base \
    musl-dev

# ENV variables
ENV CGO_ENABLED=0 GOOS=linux
    
LABEL maintainer="crlidoruk@gmail.com" \
    version="1.0" \
    description="Base image for building"