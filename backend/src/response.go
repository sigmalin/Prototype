package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type IResponse interface {
	Add(string, any)
	Error(int, string) error
	Message() error
}

var responses = make(map[string]func(w http.ResponseWriter, r *http.Request) IResponse)

func RegisterResponse(key string, response IResponse) {
	if reflect.TypeOf(response).Kind() == reflect.Ptr {
		registerResponse(key, func(w http.ResponseWriter, r *http.Request) IResponse {
			return reflect.New(reflect.ValueOf(response).Elem().Type()).Interface().(IResponse)
		})
	} else {
		registerResponse(key, func(w http.ResponseWriter, r *http.Request) IResponse {
			return reflect.New(reflect.TypeOf(response)).Elem().Interface().(IResponse)
		})
	}
}

func registerResponse(key string, response func(w http.ResponseWriter, r *http.Request) IResponse) {
	if response == nil {
		panic("Response: cannot register response with nil value")
	}

	if _, exist := responses[key]; exist {
		panic(fmt.Errorf("Response: cannot register the same response %s", key))
	}

	responses[key] = response
}
