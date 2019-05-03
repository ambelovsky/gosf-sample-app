package main

import (
	"log"

	f "github.com/ambelovsky/gosf"
)

func init() {
	RegisterRoutes() // Configure endpoint request handlers

	// Load Config Files
	f.LoadConfig("server", "server.json")
	//f.LoadConfig("server", "server-secure.json")
}

func main() {
	// Start the server
	serverConfig := f.App.Config["server"].(map[string]interface{})
	f.Startup(serverConfig)

	for k, v := range f.App.Env {
		log.Println(k + "=" + v)
	}
}
