<!-- markdownlint-disable first-line-h1 -->
<!-- markdownlint-disable html -->
<!-- markdownlint-disable no-duplicate-header -->
<div align="center">
  <img src="https://github.com/nidea1/go-gavel/blob/main/figures/logo-by-dall-e.png?raw=true" width="40%" alt="goGavel" />
</div>

## 1. Introduction

goGavel is a gRPC-based online auction platform built with Go, structured as a microservices monorepo. This project is developed as a learning exercise to explore microservices architecture, gRPC communication, and scalable system design with Go.

The goal of goGavel is to provide a real-time auction system where users can place bids, manage their auctions, and experience a high-performance distributed system. This project focuses on best practices in Go, efficient inter-service communication with gRPC, and clean architecture principles.

## 2. Key Features

- 🚀 **gRPC** - High-performance RPC framework
- 📦 **Protobuf** - Efficient data serialization
- 🐳 **Docker** - Containerized development and deployment
- 🐘 **PostgreSQL** - Robust data persistence
- 🔧 **Makefile** - Simplified build automation
- ⚙️ **Environment Variables** - Flexible configuration
- 🎯 **Clean Architecture** - Maintainable and testable code

## 3. Core Services

- **Auth** - User authentication and authorization
- **Auction** - Auction management
- **Bid** - Real-time bidding
- **Payment** - Payment processing
... and more to come!

## 4. Directory Structure

```plaintext
go-gavel-microservices/
├── build/                              # Dockerfiles for building the services 
│   ├── base/
│   │   └── Dockerfile                  # Base builder image for all services includes Go + Protobuf + gRPC
│   └── auth/
│       └── Dockerfile
├── cmd/                                # Entry point for each service
│   └── auth/
│       └── main.go
├── internal/
│   └── proto/                          # Generated protobuf files
│       ├── auth/
│       └── ...
├── proto/                              # Protobuf files
│   ├── auth.proto
│   └── ...
├── services/                           # Service implementations
│   ├── auth/
│   │   ├── api/
│   │   │   └── auth_service.go         # gRPC service implementation
│   │   ├── domain/
│   │   │   └── user.go                 # Domain model
│   │   ├── repository/
│   │   │   └── user_repository.go      # Repository implementation
│   │   └── usecase/
│   │       └── user_usecase.go         # Use case implementation
│   └── ...
├── compose/                            # Docker compose files
│   └── local/
│       └── docker-compose.local.yml
├── .envs/                              # Environment variables
│   └── .local/
│       ├── .auth
│       ├── .postgres
│       └── ...
├── scripts/                            # Helper scripts
│   ├── generate_protos.sh
│   └── ...
├── go.mod
├── go.sum
├── Makefile
├── .gitignore
├── .dockerignore
├── figures/                            # Diagrams and images for documentation
│   └── ...
├── LICENSE
└── README.md
```

## 5. Project Status

### 1. Service Implementation

- [ ] **Auth Service**
  - User authentication and authorization with JWT
  - Role-based access control
  - User profile management
- [ ] **Auction Service**
  - Auction creation and management
  - Real-time auction status updates
  - Automatic auction completion
- [ ] **Bid Service**
  - Real-time bid processing
  - Bid validation and verification
  - Historical bid tracking
- [ ] **Payment Service**
  - Secure payment processing
  - Multiple payment method support
  - Transaction history

### 2. Technical Roadmap

- [ ] **Security**
  - JWT authentication
  - Role-based access control (RBAC)
  - API rate limiting
- [ ] **Quality Assurance**
  - Unit tests
  - Integration tests
  - Load testing
- [ ] **Documentation**
  - Development guidelines

## 6. Getting Started

Documentation coming soon...

## 7. License

This code repository is licensed under the MIT License. See [LICENSE](https://github.com/nidea1/go-gavel/blob/main/LICENSE) for more information.

## 8. Contact

If you have any questions or suggestions, feel free to reach out to me at [crlidoruk@gmail.com](mailto:crlidoruk@gmail.com).
