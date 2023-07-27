package helloworld

import (
	"response"
	"response/code"
)

type Result struct {
	Message string `json:"Message" example:"hello world"`
}

func Handle(resp *response.Body) {

	resp.Code = code.SUCCESS
	resp.Message = ""
	resp.Data = &Result{Message: "hello world"}
}
