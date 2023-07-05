package content

import (
	"response"
	"session"
)

func Handle(res response.Response, ses session.Session) {
	data := make(map[string]interface{})
	data["content"] = ses.Get()

	res.Add("data", data)
	res.Message()
}
