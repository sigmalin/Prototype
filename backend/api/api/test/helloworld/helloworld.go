package helloworld

import (
	"response"
)

func Handle(res response.Response) {
	data := make(map[string]interface{})
	data["message"] = "hello world"

	res.Add("data", data)
	res.Message()
}
