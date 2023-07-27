package users

import (
	"config"
	"context"
	"net/http"
	"response"
	"response/code"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"jwtGin"
	"uid"

	db "connect/mongo"

	"api/users/login"
	"api/users/signin"
)

type handler struct {
	generator uid.Generator
	mainDB    *mongo.Database
	jwtMgr    *jwtGin.Manager
}

func newHandler() *handler {
	return &handler{
		generator: uid.GetGenerator(config.UID_GENERATOR_KEY),
		mainDB:    db.GetDB(config.DATABASE_TABLE),
		jwtMgr:    jwtGin.NewManager(config.JWT_SIGNING_KEY),
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

func (h *handler) newToken() string {
	uid, _ := h.generator()
	return uid
}

// @Summary User Signin
// @Tags users
// @version 1.0
// @produce application/json
// @Success 200 {object} response.Body{data=signin.Result} "Success"
// @Router /users/signin [get]
func (h *handler) signin(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DATABASE_TIMEOUT)
	defer cancel()

	resp := h.newResponse()

	args := signin.NewArguments(
		h.mainDB,
		ctx,
		h.newToken(),
		h.jwtMgr)

	signin.SignIn(args, resp)

	h.send(c, resp)
}

// @Summary User Login
// @Tags users
// @version 1.0
// @produce application/json
// @Param token formData string true "login token"
// @Success 200 {object} response.Body{data=login.Result} "Success"
// @Failure 400 {object} response.Body "Login Failure"
// @Router /users/login [post]
func (h *handler) login(c *gin.Context) {

	token := c.PostForm("token")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DATABASE_TIMEOUT)
	defer cancel()

	res := h.newResponse()

	args := login.NewArguments(
		h.mainDB,
		ctx,
		token,
		h.jwtMgr)

	login.LogIn(args, res)

	h.send(c, res)
}
