package models


type User struct{
	Email string `bson:"email" json:"email"`
	Role string `bson:"role" json:"role"`
	Password string `bson:"pwd" json:"-"`
}