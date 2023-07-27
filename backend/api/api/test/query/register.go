package query

import (
	"github.com/gin-gonic/gin"
)

func RegisterSubEndPoints(group *gin.RouterGroup) {

	h := newHandler()

	group.Group("/query").
		POST("allusers", h.allUsers)
}
