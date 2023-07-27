package tokenData

import "jwtGin"

type Content struct {
	JsonWebToken string `json:"JsonWebToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

func NewContent(userID string, jwtMgr *jwtGin.Manager) (*Content, error) {
	// TODO: more info
	claims := jwtGin.NewClaims(userID)
	token, err := jwtMgr.Generate(claims)
	if err != nil {
		return nil, err
	}
	return &Content{JsonWebToken: token}, nil
}
