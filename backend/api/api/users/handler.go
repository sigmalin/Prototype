package users

import (
	"config"
	"context"
	"net/http"
	"response"
	"response/code"

	"database/sql"

	"github.com/gin-gonic/gin"

	"connect/db"
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

func (h *handler) newToken() string {
	uid, _ := h.generator()
	return uid
}

func (h *handler) newSession(c *gin.Context) session.Session {
	return h.sessionMgr.SessionStart(c.Writer, c.Request)
}

// @Summary User Signin
// @Tags users
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=signin.signInData} "Success"
// @Router /users/signin [get]
func (h *handler) signin(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	res := h.newResponse()

	args := signin.NewArguments(
		h.mainDB,
		ctx,
		h.newSession(c),
		h.newToken())

	signin.SignIn(args, res)

	h.send(c, res)
}

// @Summary User Login
// @Tags users
// @version 1.0
// @produce application/json
// @Param token formData string true "login token"
// @Success 200 {object} response.Body{data=login.logInData} "Success"
// @Router /users/login [post]
func (h *handler) login(c *gin.Context) {

	token := c.PostForm("token")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	res := h.newResponse()

	args := login.NewArguments(
		h.mainDB,
		ctx,
		h.newSession(c),
		token)

	login.LogIn(args, res)

	h.send(c, res)
}
