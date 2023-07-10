package content

import (
	"response"
	"session"
)

type content struct {
	Content string `json:"content"`
}

func Handle(ses session.Session, res *response.Body) {

	res.Data = &content{Content: ses.Get().(string)}
}
