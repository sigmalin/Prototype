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

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @host 127.0.0.1:80
// @schemes http
func main() {

	initial.AllService()

	api.SSLService()
}
