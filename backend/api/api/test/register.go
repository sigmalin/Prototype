package test

import (
	"github.com/gin-gonic/gin"

	"api/test/query"
)

func RegisterEndPoints(router *gin.Engine) {

	h := newHandler()

	group := router.Group("/test")

	query.RegisterSubEndPoints(group)

	group.GET("helloworld", h.helloworld)
}
