package main

import (
	f "github.com/ambelovsky/gosf"
)

func init() {
	RegisterRoutes() // Configure endpoint request handlers

	// Load Config File Based on Environmental Configuration
	if value, exist := f.App.Env["GOSF_ENV"]; exist && value != "dev" {
		// Prod/Stage Config
		f.LoadConfig("server", "server-secure.json")
	} else {
		// Default and "dev" config
		f.LoadConfig("server", "server.json")
	}
}

func main() {
	// Start the server
	serverConfig := f.App.Config["server"].(map[string]interface{})
	f.Startup(serverConfig)
}
