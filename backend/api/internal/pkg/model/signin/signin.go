package signinData

import (
	"jwtGin"
	"model/loginData"
)

type Content struct {
	Token string `json:"Token" example:"d704e538-4f2f-486d-a2a1-a2b0ad3b4fe7"`

	Login loginData.Content `json:"login"`
}

func NewContent(token string, userID int64, jwtMgr *jwtGin.Manager) (*Content, error) {

	login, err := loginData.NewContent(userID, jwtMgr)
	if err != nil {
		return nil, err
	}

	return &Content{
		Token: token,
		Login: loginData.Content{
			Token: login.Token,
		},
	}, nil
}
