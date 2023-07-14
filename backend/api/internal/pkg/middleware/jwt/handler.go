package middleware

import (
	"config"
	"net/http"
	"response"
	"response/code"
	"strings"

	"jwtGin"

	"github.com/gin-gonic/gin"
)

type handler struct {
	jwtMgr *jwtGin.Manager
}

func NewHandler() *handler {
	return &handler{
		jwtMgr: jwtGin.NewManager(config.JWT_SIGNING_KEY),
	}
}

func (h *handler) newResponse() *response.Body {
	return &response.Body{Code: code.SUCCESS, Message: ""}
}

func (h *handler) send(c *gin.Context, res *response.Body) {
	c.JSON(http.StatusUnauthorized, res)
}

func (h *handler) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			res := h.newResponse()
			res.Error(code.AUTH_FAIURE, "Malformed Toiken")
			h.send(c, res)
			c.Abort()
			return
		} else {
			jwtToken := authHeader[1]
			claims, err := h.jwtMgr.Parse(jwtToken)
			if err != nil {
				res := h.newResponse()
				res.Error(code.AUTH_FAIURE, err.Error())
				h.send(c, res)
				c.Abort()
				return
			}
			c.Set(config.JWT_CLAIMS_KEY, claims)
			c.Next()
		}
	}
}
