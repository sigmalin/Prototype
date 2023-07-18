package modify

import (
	"github.com/gin-gonic/gin"
)

func RegisterSubEndPoints(group *gin.RouterGroup) {

	h := newHandler()

	group.Group("/modify").
		POST("bank", h.bank)
}
