package userModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	UUID     string
	Name     string   `bson:"name" json:"name"     validate:"required,min=4,max=100"`
	Nickname string   `bson:"nickname" json:"nickname" validate:"required,alphanum,min=4,max=32"`
	Birth    string   `bson:"birth" json:"birth"    validate:"required"`
  Stack    []string `bson:"stack,omitempty" json:"stack" validate:"required,dive,alpha,min=2"`

  
}