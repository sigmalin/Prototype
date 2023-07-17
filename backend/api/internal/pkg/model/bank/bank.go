package bankData

type Content struct {
	Coin     int64 `json:"Coin" example:"1000"`
	Faith    int64 `json:"Faith" example:"0"`
	Gems     int64 `json:"Gems" example:"0"`
	Treasure int64 `json:"Treasure" example:"0"`
}

func NewContent() (*Content, error) {
	return &Content{}, nil
}
