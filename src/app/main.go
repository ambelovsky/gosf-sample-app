package main

import (
	f "github.com/ambelovsky/gosf"
)

func init() {
	RegisterPlugins() // Register plugins
	RegisterRoutes()  // Configure endpoint request handlers

	// Load Config Files
	f.LoadConfig("server", "server.json")
	//f.LoadConfig("server", "server-secure.json")
}

func main() {
	// Start the server
	f.Startup(f.Config["server"].(map[string]interface{}))

	// Load Plugin App Methods
	LoadPluginMethods()
}
