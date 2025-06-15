package main

import (
	"fmt"
	"pxr-sso-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("SSO API config: %v", cfg)
}
