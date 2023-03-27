package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"custom/config"

	"custom/response"
	"custom/response/code"

	"custom/response/simple/arguments"

	"custom/session"
	_ "custom/session/redis"

	"custom/connect/db"
)

type Response = response.Response
type SessionManager = session.SessionManager

var SessionMgr *SessionManager

func apiService() {

	initSession()

	// Serve 200 status on / for k8s health checks
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/Debug", handleDebug)
	http.HandleFunc("/Query", handleQuery)
	http.HandleFunc("/Signin", handleSignin)
	http.HandleFunc("/Login", handleLogin)
	http.HandleFunc("/Session", handleSession)
	http.HandleFunc("/Form", handleForm)

	// Run the HTTP server using the bound certificate and key for TLS
	if err := http.ListenAndServe(config.API_PORT, nil); err != nil {
		log.Print("HTTP server failed to run")
	} else {
		log.Printf("HTTP server is running on port %s", config.API_PORT)
	}
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

func getArguments(w http.ResponseWriter, r *http.Request) response.Arguments {
	return arguments.SimpleResponseArguments{Writer: w, Request: r}
}

// Let / return Healthy and status code 200
func handleRoot(w http.ResponseWriter, r *http.Request) {

	res := response.GetResponse(getArguments(w, r))
	res.Message()
}

func handleQuery(w http.ResponseWriter, r *http.Request) {

	res := response.GetResponse(getArguments(w, r))

	r.ParseForm()
	name := r.PostFormValue("name")

	ctx, cancel := context.WithTimeout(r.Context(), config.SQL_TIMEOUT)
	defer cancel()

	queryUsers(db.GetDB(config.SQL_DATABASE), res, ctx, name)
}

func handleDebug(w http.ResponseWriter, r *http.Request) {

	res := response.GetResponse(getArguments(w, r))

	ctx, cancel := context.WithTimeout(r.Context(), config.SQL_TIMEOUT)
	defer cancel()

	querySleep(db.GetDB(config.SQL_DATABASE), res, ctx)
}

func handleSignin(w http.ResponseWriter, r *http.Request) {

	res := response.GetResponse(getArguments(w, r))

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		res.Error(code.UNKNOWN_ERROR, "ParseForm Failure")
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if name == "" || email == "" || password == "" {
		res.Error(code.INPUT_FAIURE, "mismatch input argument")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.SQL_TIMEOUT)
	defer cancel()

	userID, err := ExecSignin(db.GetDB(config.SQL_DATABASE), ctx, name, email, password)
	if err != nil {
		res.Error(code.SIGNIN_FAIURE, fmt.Sprintf("Cannot create user %s", email))
		return
	}

	session := SessionMgr.SessionStart(w, r)
	session.Set(userID)

	res.Message()
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	res := response.GetResponse(getArguments(w, r))

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		res.Error(code.UNKNOWN_ERROR, "ParseForm Failure")
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if email == "" || password == "" {
		res.Error(code.INPUT_FAIURE, "mismatch input argument")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), config.SQL_TIMEOUT)
	defer cancel()

	userID, err := QueryLogin(db.GetDB(config.SQL_DATABASE), ctx, email, password)
	if err != nil {
		res.Error(code.DATA_NOT_FIND, fmt.Sprintf("Cannot find user %s", email))
		return
	}

	if err := UpdateLoginTime(db.GetDB(config.SQL_DATABASE), ctx, userID); err != nil {
		log.Printf("UpdateLoginTime Failure : %s", err.Error())
	}

	session := SessionMgr.SessionStart(w, r)
	session.Set(userID)

	res.Message()
}

func handleSession(w http.ResponseWriter, r *http.Request) {

	res := response.GetResponse(getArguments(w, r))

	session, err := SessionMgr.SessionRead(w, r)
	if err != nil {
		res.Error(code.SESSION_FAIURE, err.Error())
		return
	}

	res.Add("value", session.Get())
	res.Message()
}

func handleForm(w http.ResponseWriter, r *http.Request) {

	// x-www-form-urlencoded
	/*
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
	*/

	// form-data
	if err := r.ParseMultipartForm(1024); err != nil {
		log.Print(err)
	}

	log.Print("Form: ", r.Form)
	log.Print("PostForm: ", r.PostForm)
	log.Print("FormValue: ", r.FormValue("name"))
	log.Print("PostFormValue: ", r.PostFormValue("name"))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "Test Form")
	if err != nil {
		log.Print("err")
	}
}
