package main

import (
	"github.com/loyalty-application/go-gin-backend/config"
)

// @title go-gin-backend
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.InitEnvironment()
	InitRoutes()

}
