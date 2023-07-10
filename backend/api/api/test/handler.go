package test

import (
	"github.com/gin-gonic/gin"

	rgin "response/gin"

	"api/test/helloworld"
)

type handler struct {
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) getResponse(c *gin.Context) *rgin.GinResponse {
	return rgin.NewResponse(c)
}

// @Summary Recive HelloWorld from server
// @Tags test
// @version 1.0
// @produce application/json
// @Success 200 {string} string "sigmaYAYA"
// @Router /test/helloworld [get]
func (h *handler) helloworld(c *gin.Context) {

	res := h.getResponse(c)

	helloworld.Handle(res)
}
