package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id 				bson.ObjectId	`bson:"_id,omitempty"`
	FirstName		string			`json:"first_name"`
	LastName		string			`json:"last_name"`
	Role			string			`json:"role"`
	Email			string			`json:"email"`
	PassWord		string			`json:"password,omitempty"`
	HashPassword	[]byte			`json:"hash_password,omitempty"`
}

type LoginValidate struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}
