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

func (h *handler) send(c *gin.Context, resp *response.Body) {
	if resp.Code == code.SUCCESS {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, resp)
	}
}

// @Summary Recive HelloWorld from server
// @Tags test
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=helloworld.Result} "Success"
// @Router /test/helloworld [get]
func (h *handler) helloworld(c *gin.Context) {

	resp := h.newResponse()

	helloworld.Handle(resp)

	h.send(c, resp)
}
