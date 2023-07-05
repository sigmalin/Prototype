package main

import (
	"api"
	"initial"
)

// Main will set up an http server and three endpoints
func main() {

	initial.AllService()

	api.SSLService()
}
