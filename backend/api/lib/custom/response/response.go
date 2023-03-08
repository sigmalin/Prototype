package response

import (
	"fmt"
	"net/http"
	"reflect"
)

type Response interface {
	Add(string, interface{})
	Error(int, string) error
	Message() error
}

var responses = make(map[string]func(w http.ResponseWriter, r *http.Request) Response)

func Register(key string, response Response) {
	if reflect.TypeOf(response).Kind() == reflect.Ptr {
		RegisterFactory(key, func(w http.ResponseWriter, r *http.Request) Response {
			return reflect.New(reflect.ValueOf(response).Elem().Type()).Interface().(Response)
		})
	} else {
		RegisterFactory(key, func(w http.ResponseWriter, r *http.Request) Response {
			return reflect.New(reflect.TypeOf(response)).Elem().Interface().(Response)
		})
	}
}

func RegisterFactory(key string, response func(w http.ResponseWriter, r *http.Request) Response) {
	if response == nil {
		panic("Response: cannot register response with nil value")
	}

	if _, exist := responses[key]; exist {
		panic(fmt.Errorf("Response: cannot register the same response %s", key))
	}

	responses[key] = response
}

func GetFactory(key string) func(w http.ResponseWriter, r *http.Request) Response {
	return responses[key]
}
