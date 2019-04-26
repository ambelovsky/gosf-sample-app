package main

import f "github.com/ambelovsky/gosf"

// Util contains all util.* endpoint controllers
type Util struct{}

// Echo returns the passed message back to the client
func (controller Util) Echo(request *f.Request, clientMessage *f.Message) {
	response := new(f.Message)
	response.Success = true
	response.Message = PluginSamplePlugin.Echo(clientMessage.Message)

	request.Respond(clientMessage, response)
}
