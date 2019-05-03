package main

import f "github.com/ambelovsky/gosf"

// Util contains all util.* endpoint controllers
type Util struct{}

// Echo returns the passed message back to the client
func (controller Util) Echo(client *f.Client, request *f.Request) *f.Message {
	var description string

	// Argument Parsing
	if val, ok := request.Message.Body["description"]; ok {
		description = val.(string)
	}

	response := new(f.Message)
	response.Success = true
	response.Text = PluginSamplePlugin.Echo(request.Message.Text)

	// If a detailed description was entered, send it back to the client
	if description != "" {
		response.Text += " - " + description
	}

	return response
}
