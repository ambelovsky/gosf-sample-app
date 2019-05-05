package main

import (
	"github.com/ambelovsky/gosf"
	f "github.com/ambelovsky/gosf"
)

// Auth contains all auth.* endpoint controllers
type Auth struct{}

// Register is generally the first endpoint requested by a connected client
func (controller Auth) Register(client *f.Client, request *f.Request) *f.Message {
	response := f.NewSuccessMessage("client registered", nil)

	adminMessage := f.NewSuccessMessage("new client registered", nil)
	gosf.Broadcast("admin", "admin.register", adminMessage)

	return response
}
