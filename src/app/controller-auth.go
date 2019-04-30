package main

import f "github.com/ambelovsky/gosf"

// Auth contains all auth.* endpoint controllers
type Auth struct{}

// Register is generally the first endpoint requested by a connected client
func (controller Auth) Register(client *f.Client, request *f.Request) *f.Message {
	response := new(f.Message)
	response.Success = true
	response.Text = "client registered"

	return response
}
