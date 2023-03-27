package simple

import (
	"custom/response/code"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type any interface{}
type dict map[string]any

type SimpleResponse struct {
	writer  http.ResponseWriter
	request *http.Request
	data    dict
}

func (r SimpleResponse) send(msg string) {
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(http.StatusOK)
	_, err := io.WriteString(r.writer, msg)
	if err != nil {
		log.Print(err)
	}
}

func (r SimpleResponse) Add(key string, value interface{}) {
	r.data[key] = value
}

func (r SimpleResponse) Error(code int, msg string) error {
	res := dict{"error": dict{"code": code, "message": msg}}

	data, err := json.Marshal(res)
	if err != nil {
		log.Print(err.Error())
	} else {
		r.send(string(data))
	}
	return err
}

func (r SimpleResponse) Message() error {
	r.data["error"] = dict{"code": code.SUCCESS, "message": ""}

	data, err := json.Marshal(r.data)
	if err != nil {
		log.Print(err.Error())
		r.Error(code.UNKNOWN_ERROR, err.Error())
	} else {
		r.send(string(data))
	}
	return err
}

func NewResponse(w http.ResponseWriter, r *http.Request) SimpleResponse {
	return SimpleResponse{writer: w, request: r, data: dict{}}
}
