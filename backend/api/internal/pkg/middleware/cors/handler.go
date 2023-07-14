package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 由 nginx 處理
		//c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		//c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		//c.Header("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, PUT, OPTIONS")
		//c.Header("Access-Control-Allow-Credentials", "true")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		// preflight 對應
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
