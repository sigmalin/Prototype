package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.Engine) {

	h := newHandler()

	router.Group("/users").
		GET("/signin", h.signin).
		POST("/login", h.login)
}
