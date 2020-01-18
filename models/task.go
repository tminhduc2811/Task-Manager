package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	Id			bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	CreatedBy	string			`json:"created_by"`
	Name		string			`json:"name"`
	Description	string			`json:"description"`
	CreatedOn	time.Time		`json:"created_on"`
	Due			time.Time		`json:"due"`
	Status		string			`json:"status"`
	Tags		[]string		`json:"tags"`
}

