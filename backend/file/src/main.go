package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	FILE_PORT         string
	SSL_CERTIFICATION string
	SSL_PRIVATE_KEY   string
)

// Main will set up an http server and three endpoints
func main() {

	initEnv()

	router := gin.Default()

	router.Static("/bundles", "./bundles")

	if err := router.RunTLS(FILE_PORT, SSL_CERTIFICATION, SSL_PRIVATE_KEY); err != nil {
		log.Print("HTTP server failed to run")
	} else {
		log.Printf("HTTP server is running on port %s, msg = %s", FILE_PORT, err)
	}
}

func initEnv() {
	FILE_PORT = os.Getenv("FILE_PORT")
	SSL_CERTIFICATION = os.Getenv("SSL_CERTIFICATION")
	SSL_PRIVATE_KEY = os.Getenv("SSL_PRIVATE_KEY")
}
