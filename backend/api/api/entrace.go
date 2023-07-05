package api

import (
	"github.com/gin-gonic/gin"

	"api/test"
	"api/users"
)

func entrace(router *gin.Engine) {

	test.RegisterEndPoints(router)
	users.RegisterEndPoints(router)
}
