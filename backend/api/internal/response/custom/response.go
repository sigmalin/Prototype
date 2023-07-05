package custom

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"response/code"
)

type any interface{}
type dict map[string]any

type CustomResponse struct {
	writer  http.ResponseWriter
	request *http.Request
	data    dict
}

func (r CustomResponse) send(msg string) {
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(http.StatusOK)
	_, err := io.WriteString(r.writer, msg)
	if err != nil {
		log.Print(err)
	}
}

func (r CustomResponse) Add(key string, value interface{}) {
	r.data[key] = value
}

func (r CustomResponse) Error(code int, msg string) error {
	res := dict{"error": dict{"code": code, "message": msg}}

	data, err := json.Marshal(res)
	if err != nil {
		log.Print(err.Error())
	} else {
		r.send(string(data))
	}
	return err
}

func (r CustomResponse) Message() error {
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
