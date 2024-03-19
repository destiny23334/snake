package main

import (
	"snake/global"
	"snake/router"
)

func main() {
	// Initialize Log
	global.Logger.Info("Logger initialized")

	// Load the Config
	// TODO: 从配置中心读
	global.LoadConfig()

	// Initialize Database

	// Initialize Middleware

	// Start server
	server := router.Routers()
	err := server.Run(":11001")
	if err != nil {
		return
	}
}
