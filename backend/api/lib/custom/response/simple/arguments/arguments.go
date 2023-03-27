package arguments

import (
	"net/http"

	"custom/response"
	"custom/response/simple"
)

const typeName = "simple"

type SimpleResponseArguments struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (args SimpleResponseArguments) Type() string {
	return typeName
}

func GetSimpleResponse(args response.Arguments) response.Response {
	data, ok := args.(SimpleResponseArguments)
	if !ok {
		return nil
	}

	return simple.NewResponse(data.Writer, data.Request)
}

// init
func init() {
	response.Register(typeName, GetSimpleResponse)
}
