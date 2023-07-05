package api

import (
	"config"
	"log"

	"github.com/gin-gonic/gin"
)

func SSLService() {
	engine := gin.Default()

	entrace(engine)

	if err := engine.RunTLS(config.API_PORT, config.SSL_CERTIFICATION, config.SSL_PRIVATE_KEY); err != nil {
		log.Print("HTTP server failed to run")
	} else {
		log.Printf("HTTP server is running on port %s, msg = %s", config.API_PORT, err)
	}
}
