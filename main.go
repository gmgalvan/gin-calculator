package main

import (
	"github.com/gin-calculator/routes"
)

// @title Calculator API
// @description Calculator server API

// @host      localhost:8080
// @BasePath /calculator/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	server := routes.NewRouter()
	server.Run()
}
