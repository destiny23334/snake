package main

import (
	"snake/pkg"
	"snake/router"
)

func main() {
	// Initialize Log
	pkg.Logger.Info("Logger initialized")

	// Initialize Register Center

	// Load the Config

	// Initialize Database

	// Initialize Middleware

	// Start server
	server := router.Routers()
	err := server.Run(":11001")
	if err != nil {
		return
	}
}
