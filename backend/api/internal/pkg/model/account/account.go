package accountData

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {
	Token      string `json:"Token" bson:"Token" example:"d704e538-4f2f-486d-a2a1-a2b0ad3b4fe7"`
	Mail       string `json:"Mail" bson:"Mail,omitempty" example:"sigma@company.com"`
	CreateTime int64  `json:"CreateTime" bson:"CreateTime" example:"1690361379"`
	UpdateTime int64  `json:"UpdateTime" bson:"UpdateTime" example:"1690361379"`
}

type ID struct {
	ID primitive.ObjectID `json:"-" bson:"_id"`
}
