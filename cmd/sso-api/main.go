package main

import (
	"pxr-sso-api/internal/app"
)

// @title SSO API
// @version 1.0
// @description API server for SSO

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	application := app.New()

	go application.Start()

	application.GracefulStop()
}
