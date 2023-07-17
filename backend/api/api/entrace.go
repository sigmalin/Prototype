package api

import (
	"github.com/gin-gonic/gin"

	middleware "middleware/cors"

	"api/me"
	"api/test"
	"api/users"
)

func entrace(router *gin.Engine) {

	cors := middleware.NewHandler()

	router.Use(cors.Handle())

	test.RegisterEndPoints(router)
	users.RegisterEndPoints(router)
	me.RegisterEndPoints(router)
}
