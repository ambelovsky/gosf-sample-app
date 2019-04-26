package main

import f "github.com/ambelovsky/gosf"

func RegisterRoutes() {
	auth := new(Auth)
	f.Listen("auth.register", auth.Register)
}
