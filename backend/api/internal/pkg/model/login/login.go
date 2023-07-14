package loginData

import (
	"strconv"

	"jwtGin"
)

type Content struct {
	Token string `json:"JWT" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

func NewContent(userID int64, jwtMgr *jwtGin.Manager) (*Content, error) {
	// TODO: more info
	claims := jwtGin.NewClaims(strconv.FormatInt(userID, 10))
	token, err := jwtMgr.Generate(claims)
	if err != nil {
		return nil, err
	}
	return &Content{Token: token}, nil
}
