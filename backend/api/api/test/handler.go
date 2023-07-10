package test

import (
	"net/http"
	"response"
	"response/code"

	"github.com/gin-gonic/gin"

	"api/test/helloworld"
)

type handler struct {
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) newResponse() *response.Body {
	return &response.Body{Code: code.SUCCESS, Message: ""}
}

func (h *handler) send(c *gin.Context, res *response.Body) {
	if res.Code == code.SUCCESS {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, res)
	}
}

// @Summary Recive HelloWorld from server
// @Tags test
// @version 1.0
// @produce application/json
// @Success 200 {string} string "sigmaYAYA"
// @Router /test/helloworld [get]
func (h *handler) helloworld(c *gin.Context) {

	res := h.newResponse()

	helloworld.Handle(res)

	h.send(c, res)
}
