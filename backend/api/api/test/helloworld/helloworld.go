package helloworld

import (
	"response"
	"response/code"
)

type content struct {
	Message string `json:"message"`
}

func Handle(res *response.Body) {

	res.Code = code.SUCCESS
	res.Message = ""
	res.Data = &content{Message: "hello world"}
}
