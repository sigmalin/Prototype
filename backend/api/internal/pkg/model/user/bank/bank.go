package bankData

type Content struct {
	Coin     int64 `json:"Coin" bson:"Coin" example:"1000"`
	Faith    int64 `json:"Faith" bson:"Faith" example:"0"`
	Gems     int64 `json:"Gems" bson:"Gems" example:"0"`
	Treasure int64 `json:"Treasure" bson:"Treasure" example:"0"`
}

func NewContent() (*Content, error) {
	return &Content{
		Coin:     1000,
		Faith:    0,
		Gems:     0,
		Treasure: 0,
	}, nil
}
