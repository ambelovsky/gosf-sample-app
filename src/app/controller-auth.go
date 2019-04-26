package main

import f "github.com/ambelovsky/gosf"

// Auth contains all auth.* endpoint controllers
type Auth struct{}

// Register is generally the first endpoint requested by a connected client
func (controller Auth) Register(request *f.Request, clientMessage *f.Message) {
	response := new(f.Message)
	response.Success = true
	response.Message = "client registered"

	request.Respond(clientMessage, response)
}
