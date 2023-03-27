package response

import (
	"fmt"
)

type Response interface {
	Add(string, interface{})
	Error(int, string) error
	Message() error
}

var responses = make(map[string]func(args Arguments) Response)

func Register(key string, response func(args Arguments) Response) {
	if response == nil {
		panic("Response: cannot register response with nil value")
	}

	if _, exist := responses[key]; exist {
		panic(fmt.Errorf("Response: cannot register the same response %s", key))
	}

	responses[key] = response
}

func GetFactory(key string) func(args Arguments) Response {
	return responses[key]
}

func GetResponse(args Arguments) Response {
	processor, exist := responses[args.Type()]
	if !exist {
		panic(fmt.Errorf("Response: type %s isnot exist", args.Type()))
	}
	return processor(args)
}
