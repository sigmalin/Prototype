package me

import (
	"api/me/bank"
	"config"
	"context"
	"net/http"
	"response"
	"response/code"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	db "connect/mongo"
	"jwtGin"
)

type handler struct {
	mainDB *mongo.Database
	jwtMgr *jwtGin.Manager
}

func newHandler() *handler {
	return &handler{
		mainDB: db.GetDB(config.DATABASE_TABLE),
		jwtMgr: jwtGin.NewManager(config.JWT_SIGNING_KEY),
	}
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

// @Summary Query My Bank Data
// @Tags me
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=bank.Result} "Success"
// @Router /me/bank [get]
// @Security Bearer
func (h *handler) bank(c *gin.Context) {

	resp := h.newResponse()

	claims, exist := c.Get(config.JWT_CLAIMS_KEY)
	if !exist {
		resp.Error(code.SESSION_FAIURE, "session failure")
		h.send(c, resp)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DATABASE_TIMEOUT)
	defer cancel()

	args := bank.NewArguments(
		h.mainDB,
		ctx,
		claims.(*jwtGin.Claims).ID,
	)

	bank.Bank(args, resp)

	h.send(c, resp)
}
