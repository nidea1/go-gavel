module github.com/nidea1/go-gavel/services/auth

go 1.24.0

require (
    github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/nidea1/go-gavel/pkg v0.0.1
	github.com/nidea1/go-gavel/proto v0.0.1
)

replace github.com/nidea1/go-gavel/pkg => ../../pkg

replace github.com/nidea1/go-gavel/proto => ../../proto
