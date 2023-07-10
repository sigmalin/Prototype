package query

import (
	"config"
	"net/http"
	"response"
	"response/code"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"
	"session"

	"api/test/query/allusers"
	content "api/test/query/sessioncontent"
)

type handler struct {
	sessionMgr *session.SessionManager
	mainDB     *sql.DB
}

func newHandler() *handler {
	return &handler{
		sessionMgr: session.GetManager(config.SESSION_MANAGER_KEY),
		mainDB:     db.GetDB(config.SQL_DATABASE),
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

func (h *handler) getSession(c *gin.Context) (session.Session, error) {
	return h.sessionMgr.SessionRead(c.Writer, c.Request)
}

// @Summary Query Session Content
// @Tags test
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=content.content} "Success"
// @Failure 400 {object} response.Body "Failure"
// @Router /test/query/session [get]
func (h *handler) sessionContent(c *gin.Context) {

	res := h.newResponse()

	ses, err := h.getSession(c)
	if err != nil {
		res.Error(code.SESSION_FAIURE, err.Error())
		h.send(c, res)
		return
	}

	content.Handle(ses, res)

	h.send(c, res)
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
