package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"custom/config"

	"custom/response"
	"custom/response/code"

	"custom/response/gin/arguments"

	"custom/session"
	_ "custom/session/redis"

	"custom/connect/db"
)

type Response = response.Response
type SessionManager = session.SessionManager

var SessionMgr *SessionManager

func apiService() {

	initSession()

	initSSLService()
}

func initSession() {

	const providerType = "redis"

	mgr, err := session.CreateSessionManager(providerType, config.SESSION_NAME)
	if err != nil {
		log.Fatal(fmt.Errorf("cannot get session : %s", providerType))
	}

	SessionMgr = mgr
	go SessionMgr.GC(config.SESSION_EXPIRATION)
}

func initService() {
	engine := gin.Default()

	engine.GET("/", handleRoot)
	engine.GET("/Debug", handleDebug)

	engine.POST("/Query", handleQuery)
	engine.POST("/Signin", handleSignin)
	engine.POST("/Login", handleLogin)
	engine.POST("/Session", handleSession)

	if err := engine.Run(config.API_PORT); err != nil {
		log.Print("HTTP server failed to run")
	} else {
		log.Printf("HTTP server is running on port %s", config.API_PORT)
	}
}

func initSSLService() {
	engine := gin.Default()

	engine.GET("/", handleRoot)
	engine.GET("/Debug", handleDebug)

	engine.POST("/Query", handleQuery)
	engine.POST("/Signin", handleSignin)
	engine.POST("/Login", handleLogin)
	engine.POST("/Session", handleSession)

	if err := engine.RunTLS(config.API_PORT, config.SSL_CERTIFICATION, config.SSL_PRIVATE_KEY); err != nil {
		log.Print("HTTP server failed to run")
	} else {
		log.Printf("HTTP server is running on port %s, msg = %s", config.API_PORT, err)
	}
}

func getArguments(c *gin.Context) response.Arguments {
	return arguments.GinResponseArguments{Context: c}
}

// Let / return Healthy and status code 200
func handleRoot(c *gin.Context) {

	res := response.GetResponse(getArguments(c))
	res.Message()
}

func handleQuery(c *gin.Context) {

	res := response.GetResponse(getArguments(c))

	name := c.PostForm("name")

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	queryUsers(db.GetDB(config.SQL_DATABASE), res, ctx, name)
}

func handleDebug(c *gin.Context) {

	res := response.GetResponse(getArguments(c))

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	querySleep(db.GetDB(config.SQL_DATABASE), res, ctx)
}

func handleSignin(c *gin.Context) {

	res := response.GetResponse(getArguments(c))

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	if name == "" || email == "" || password == "" {
		res.Error(code.INPUT_FAIURE, "mismatch input argument")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	userID, err := ExecSignin(db.GetDB(config.SQL_DATABASE), ctx, name, email, password)
	if err != nil {
		res.Error(code.SIGNIN_FAIURE, fmt.Sprintf("Cannot create user %s", email))
		return
	}

	session := SessionMgr.SessionStart(c.Writer, c.Request)
	session.Set(userID)

	res.Message()
}

func handleLogin(c *gin.Context) {

	res := response.GetResponse(getArguments(c))

	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		res.Error(code.INPUT_FAIURE, "mismatch input argument")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), config.SQL_TIMEOUT)
	defer cancel()

	userID, err := QueryLogin(db.GetDB(config.SQL_DATABASE), ctx, email, password)
	if err != nil {
		res.Error(code.DATA_NOT_FIND, fmt.Sprintf("Cannot find user %s", email))
		return
	}

	if err := UpdateLoginTime(db.GetDB(config.SQL_DATABASE), ctx, userID); err != nil {
		log.Printf("UpdateLoginTime Failure : %s", err.Error())
	}

	session := SessionMgr.SessionStart(c.Writer, c.Request)
	session.Set(userID)

	res.Message()
}

func handleSession(c *gin.Context) {

	res := response.GetResponse(getArguments(c))

	session, err := SessionMgr.SessionRead(c.Writer, c.Request)
	if err != nil {
		res.Error(code.SESSION_FAIURE, err.Error())
		return
	}

	res.Add("value", session.Get())
	res.Message()
}
