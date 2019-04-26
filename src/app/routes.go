package main

import f "gosf"

func RegisterRoutes() {
	auth := new(Auth)
	f.Listen("auth.register", auth.Register)
}
