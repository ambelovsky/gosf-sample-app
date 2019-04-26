package main

import f "github.com/ambelovsky/gosf"

// RegisterRoutes maps controller functions to endpoints
func RegisterRoutes() {
	auth := new(Auth)
	f.Listen("auth.register", auth.Register)

	util := new(Util)
	f.Listen("util.echo", util.Echo)
}
