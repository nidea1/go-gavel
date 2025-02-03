package main

import (
	"fmt"

	"github.com/nidea1/go-gavel/services/auth/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	fmt.Println("Starting server on port", cfg.Server.Port)
}
