package userModel

type User struct {
  Name     string   `bson:"name" json:"name"     validate:"required,alpha,min=4,max=15"`
  Nickname string   `bson:"nickname" json:"nickname" validate:"required,alphanum,min=4,max=15"`
  Birth    string   `bson:"birth" json:"birth"    validate:"required"`
  Stack    []string `bson:"stack,omitempty" json:"stack"    validate:"max=50"`
}
