package main

import f "github.com/ambelovsky/gosf"

// Util contains all util.* endpoint controllers
type Util struct{}

// Echo returns the passed message back to the client
func (controller Util) Echo(request *f.Request) {
	response := new(f.Message)
	response.Success = true
	response.Text = PluginSamplePlugin.Echo(request.Message.Text)

	request.Respond(response)
}
