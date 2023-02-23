package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

var GetResponse func(w http.ResponseWriter, r *http.Request) IResponse
var SessionMgr *SessionManager

func apiService() {

	initResponse()
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
	if err := http.ListenAndServe(API_PORT, nil); err != nil {
		log.Print("HTTP server failed to run")
	} else {
		log.Printf("HTTP server is running on port %s", API_PORT)
	}
}

func initResponse() {

	const responseType = "simple"

	GetResponse = responses[responseType]
	if GetResponse == nil {
		log.Fatal(fmt.Errorf("cannot get response : %s", responseType))
	}
}

func initSession() {

	const providerType = "redis"

	mgr, err := CreateSessionManager(providerType, SESSION_NAME)
	if err != nil {
		log.Fatal(fmt.Errorf("cannot get session : %s", providerType))
	}

	SessionMgr = mgr
	go SessionMgr.GC()
}

// Let / return Healthy and status code 200
func handleRoot(w http.ResponseWriter, r *http.Request) {

	res := GetResponse(w, r)
	res.Message()
}

func handleQuery(w http.ResponseWriter, r *http.Request) {

	res := GetResponse(w, r)

	r.ParseForm()
	name := r.PostFormValue("name")

	ctx, cancel := context.WithTimeout(r.Context(), SQL_TIMEOUT)
	defer cancel()

	queryUsers(MainDB, res, ctx, name)
}

func handleDebug(w http.ResponseWriter, r *http.Request) {

	res := GetResponse(w, r)

	ctx, cancel := context.WithTimeout(r.Context(), SQL_TIMEOUT)
	defer cancel()

	querySleep(MainDB, res, ctx)
}

func handleSignin(w http.ResponseWriter, r *http.Request) {

	res := GetResponse(w, r)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		res.Error(UNKNOWN_ERROR, "ParseForm Failure")
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if name == "" || email == "" || password == "" {
		res.Error(INPUT_FAIURE, "mismatch input argument")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), SQL_TIMEOUT)
	defer cancel()

	userID, err := ExecSignin(MainDB, ctx, name, email, password)
	if err != nil {
		res.Error(SIGNIN_FAIURE, fmt.Sprintf("Cannot create user %s", email))
		return
	}

	session := SessionMgr.SessionStart(w, r)
	session.Set(userID)

	res.Message()
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	res := GetResponse(w, r)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
		res.Error(UNKNOWN_ERROR, "ParseForm Failure")
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if email == "" || password == "" {
		res.Error(INPUT_FAIURE, "mismatch input argument")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), SQL_TIMEOUT)
	defer cancel()

	userID, err := QueryLogin(MainDB, ctx, email, password)
	if err != nil {
		res.Error(DATA_NOT_FIND, fmt.Sprintf("Cannot find user %s", email))
		return
	}

	if err := UpdateLoginTime(MainDB, ctx, userID); err != nil {
		log.Printf("UpdateLoginTime Failure : %s", err.Error())
	}

	session := SessionMgr.SessionStart(w, r)
	session.Set(userID)

	res.Message()
}

func handleSession(w http.ResponseWriter, r *http.Request) {

	res := GetResponse(w, r)

	session, err := SessionMgr.SessionRead(w, r)
	if err != nil {
		res.Error(SESSION_FAIURE, err.Error())
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
