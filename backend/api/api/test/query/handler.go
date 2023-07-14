package query

import (
	"config"
	"net/http"
	"response"
	"response/code"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"

	"api/test/query/allusers"
)

type handler struct {
	mainDB *sql.DB
}

func newHandler() *handler {
	return &handler{
		mainDB: db.GetDB(config.SQL_DATABASE),
	}
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

// @Summary Query All users
// @Tags test
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=[]allusers.user} "Success"
// @Router /test/query/allusers [get]
func (h *handler) allUsers(c *gin.Context) {

	res := h.newResponse()

	allusers.Handle(c.Request.Context(), h.mainDB, res)

	h.send(c, res)
}
