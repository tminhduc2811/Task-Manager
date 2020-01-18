package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id 				bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	FirstName		string			`json:"first_name"`
	LastName		string			`json:"last_name"`
	Email			string			`json:"email"`
	PassWord		string			`json:"pass_word,omitempty"`
	HashPassword	[]byte			`json:"hash_password,omitempty"`
}
