package main

import (
	"api"
	"initial"
)

// @title Prototype Api Server
// @version 1.0
// @description Standard Api Server

// @contact.name sigma
// @contact.url https://github.com/sigmalin/Prototype

// @host 127.0.0.1:443
// @schemes https
func main() {

	initial.AllService()

	api.SSLService()
}
