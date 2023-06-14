package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Neflix struct {
	//this is the _ID  //it provides the uniqueness
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
