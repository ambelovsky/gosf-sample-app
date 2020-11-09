package main

import (
	f "github.com/ambelovsky/gosf"
	samplePlugin "github.com/ambelovsky/gosf-sample-plugin"
)

// Util contains all util.* endpoint controllers
type Util struct{}

type EchoDetail struct {
	OneThing     string `json:"oneThing,omitempty"`
	AnotherThing struct {
		MoreDetail string `json:"moreDetail,omitempty"`
	} `json:"anotherThing,omitempty"`
}

type EchoRequestBody struct {
	Description string `json:"description,omitempty"`
}

// Echo returns the passed message back to the client
func (controller Util) Echo(client *f.Client, request *f.Request) *f.Message {
	// Get request arguments and convert them to a predefined struct
	requestBody := new(EchoRequestBody)
	f.MapToStruct(request.Message.Body, requestBody)

	// Use a plugin-exposed function and set the response text
	responseText := samplePlugin.Echo(request.Message.Text)

	// If a detailed description was entered, send it back to the client
	if requestBody.Description != "" {
		responseText += " - " + requestBody.Description
	}

	echoDetail := &EchoDetail{
		OneThing: "this is one thing",
	}
	echoDetail.AnotherThing.MoreDetail = "and another thing..."

	return f.NewSuccessMessage(responseText, f.StructToMap(echoDetail))
}

// GroupEcho returns the passed message back to all connected users joined to the main group except the originating client
func (controller Util) GroupEcho(client *f.Client, request *f.Request) *f.Message {
	// Get request arguments and convert them to a predefined struct
	requestBody := new(EchoRequestBody)
	f.MapToStruct(request.Message.Body, requestBody)

	// Use a plugin-exposed function and set the response text
	responseText := samplePlugin.Echo(request.Message.Text)

	// If a detailed description was entered, send it back to the client
	if requestBody.Description != "" {
		responseText += " - " + requestBody.Description
	}

	echoDetail := &EchoDetail{
		OneThing: "this is one thing",
	}
	echoDetail.AnotherThing.MoreDetail = "and another thing..."

	client.Broadcast("testRoom", request.Endpoint, f.NewSuccessMessage(responseText, f.StructToMap(echoDetail)))
	return f.NewSuccessMessage()
}
