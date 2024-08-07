package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct{
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
	IsAdmin bool `bson:"role" json:"role"`
	Password string `bson:"pwd" json:"-"`
}