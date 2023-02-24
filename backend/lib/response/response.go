package response

import (
	"fmt"
	"net/http"
	"reflect"
)

type Response interface {
	Add(string, any)
	Error(int, string) error
	Message() error
}

var responses = make(map[string]func(w http.ResponseWriter, r *http.Request) Response)

func RegisterResponse(key string, response Response) {
	if reflect.TypeOf(response).Kind() == reflect.Ptr {
		registerResponse(key, func(w http.ResponseWriter, r *http.Request) Response {
			return reflect.New(reflect.ValueOf(response).Elem().Type()).Interface().(Response)
		})
	} else {
		registerResponse(key, func(w http.ResponseWriter, r *http.Request) Response {
			return reflect.New(reflect.TypeOf(response)).Elem().Interface().(Response)
		})
	}
}

func registerResponse(key string, response func(w http.ResponseWriter, r *http.Request) Response) {
	if response == nil {
		panic("Response: cannot register response with nil value")
	}

	if _, exist := responses[key]; exist {
		panic(fmt.Errorf("Response: cannot register the same response %s", key))
	}

	responses[key] = response
}
