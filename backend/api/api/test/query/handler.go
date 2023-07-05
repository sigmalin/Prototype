package query

import (
	"config"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"
	"response/code"
	rgin "response/gin"
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

func (h *handler) getResponse(c *gin.Context) *rgin.GinResponse {
	return rgin.NewResponse(c)
}

func (h *handler) getSession(c *gin.Context) (session.Session, error) {
	return h.sessionMgr.SessionRead(c.Writer, c.Request)
}

func (h *handler) sessionContent(c *gin.Context) {

	res := h.getResponse(c)

	ses, err := h.getSession(c)
	if err != nil {
		res.Error(code.SESSION_FAIURE, err.Error())
		return
	}

	content.Handle(res, ses)
}

func (h *handler) allUsers(c *gin.Context) {

	res := h.getResponse(c)

	allusers.Handle(c.Request.Context(), res, h.mainDB)
}
