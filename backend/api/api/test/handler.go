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

func (h *handler) helloworld(c *gin.Context) {

	res := h.getResponse(c)

	helloworld.Handle(res)
}
