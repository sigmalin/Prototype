package me

import (
	"api/me/bank"
	"config"
	"context"
	"net/http"
	"response"
	"response/code"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"
	"jwtGin"
)

type handler struct {
	mainDB *sql.DB
	jwtMgr *jwtGin.Manager
}

func newHandler() *handler {
	return &handler{
		mainDB: db.GetDB(config.SQL_DATABASE),
		jwtMgr: jwtGin.NewManager(config.JWT_SIGNING_KEY),
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

// @Summary Query My Bank Data
// @Tags me
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=bankData.Content} "Success"
// @Router /me/bank [get]
// @Security Bearer
func (h *handler) bank(c *gin.Context) {

	res := h.newResponse()

	claims, exist := c.Get(config.JWT_CLAIMS_KEY)
	if !exist {
		res.Error(code.SESSION_FAIURE, "session failure")
		h.send(c, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	args := bank.NewArguments(
		h.mainDB,
		ctx,
		claims.(*jwtGin.Claims).ID,
	)

	bank.Bank(args, res)

	h.send(c, res)
}
