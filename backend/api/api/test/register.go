package test

import (
	"github.com/gin-gonic/gin"

	"api/test/modify"
	"api/test/query"
)

func RegisterEndPoints(router *gin.Engine) {

	h := newHandler()

	group := router.Group("/test")

	query.RegisterSubEndPoints(group)

	modify.RegisterSubEndPoints(group)

	group.GET("helloworld", h.helloworld)
}
