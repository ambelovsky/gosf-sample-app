package main

import f "github.com/ambelovsky/gosf"

type Auth struct{}

func (controller Auth) Register(request *f.Request, clientMessage *f.Message) {
	response := new(f.Message)
	response.Status = "success"
	response.Message = "client registered"

	// Use a sample plugin method
	PluginMessageCounter.Tick()

	request.Respond(clientMessage, response)
}
