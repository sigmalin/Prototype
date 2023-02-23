package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Response struct {
	writer  http.ResponseWriter
	request *http.Request
	data    dict
}

func (r Response) send(msg string) {
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(http.StatusOK)
	_, err := io.WriteString(r.writer, msg)
	if err != nil {
		log.Print(err)
	}
}

func (r Response) Add(key string, value any) {
	r.data[key] = value
}

func (r Response) Error(code int, msg string) error {
	res := dict{"error": dict{"code": code, "message": msg}}

	data, err := json.Marshal(res)
	if err != nil {
		log.Print(err.Error())
	} else {
		r.send(string(data))
	}
	return err
}

func (r Response) Message() error {
	r.data["error"] = dict{"code": SUCCESS, "message": ""}

	data, err := json.Marshal(r.data)
	if err != nil {
		log.Print(err.Error())
		r.Error(UNKNOWN_ERROR, err.Error())
	} else {
		r.send(string(data))
	}
	return err
}

func GetSimpleResponse(w http.ResponseWriter, r *http.Request) IResponse {
	return Response{writer: w, request: r, data: dict{}}
}

// init
func initResponseProvider() {
	registerResponse("simple", GetSimpleResponse)
}
