package arguments

import (
	"github.com/gin-gonic/gin"

	"custom/response"
	gresponse "custom/response/gin"
)

const typeName = "gin"

type GinResponseArguments struct {
	Context *gin.Context
}

func (args GinResponseArguments) Type() string {
	return typeName
}

func GetGinResponse(args response.Arguments) response.Response {
	data, ok := args.(GinResponseArguments)
	if !ok {
		return nil
	}

	return gresponse.NewResponse(data.Context)
}

// init
func init() {
	response.Register(typeName, GetGinResponse)
}
