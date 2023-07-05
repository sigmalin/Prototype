package rgin

import (
	"github.com/gin-gonic/gin"
)

func NewResponse(c *gin.Context) *GinResponse {
	return &GinResponse{context: c, data: gin.H{}}
}
