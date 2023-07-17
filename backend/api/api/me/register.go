package me

import (
	middleware "middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.Engine) {

	jwt := middleware.NewHandler()

	h := newHandler()

	router.Group("/me").
		Use(jwt.Auth()).
		GET("/bank", h.bank)
}
