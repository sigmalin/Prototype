package users

import (
	"config"
	"context"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"
	rgin "response/gin"
	"session"
	"uid"

	"api/users/login"
	"api/users/signin"
)

type handler struct {
	generator  uid.Generator
	sessionMgr *session.SessionManager
	mainDB     *sql.DB
}

func newHandler() *handler {
	return &handler{
		generator:  uid.GetGenerator(config.UID_GENERATOR_KEY),
		sessionMgr: session.GetManager(config.SESSION_MANAGER_KEY),
		mainDB:     db.GetDB(config.SQL_DATABASE),
	}
}

func (h *handler) newToken() string {
	uid, _ := h.generator()
	return uid
}

func (h *handler) getResponse(c *gin.Context) *rgin.GinResponse {
	return rgin.NewResponse(c)
}

func (h *handler) newSession(c *gin.Context) session.Session {
	return h.sessionMgr.SessionStart(c.Writer, c.Request)
}

func (h *handler) signin(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	args := signin.NewArguments(
		h.mainDB,
		ctx,
		h.getResponse(c),
		h.newSession(c),
		h.newToken())

	signin.SignIn(args)
}

func (h *handler) login(c *gin.Context) {

	token := c.PostForm("token")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	args := login.NewArguments(
		h.mainDB,
		ctx,
		h.getResponse(c),
		h.newSession(c),
		token)

	login.LogIn(args)
}
