package main

import f "github.com/ambelovsky/gosf"

func init() {
	RegisterPlugins() // Register plugins
	RegisterRoutes()  // Configure endpoint request handlers
}

func main() {
	config := map[string]interface{}{
		"port": 9999,
		"path": "/"}

	// Start the server
	f.Startup(config)

	// Load Plugin App Methods
	LoadPluginMethods()
}
