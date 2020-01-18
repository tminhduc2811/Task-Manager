package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Note struct {
	Id			bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	TaskId		bson.ObjectId	`json:"task_id"`
	Description	string			`json:"description"`
	CreatedOn	time.Time		`json:"created_on"`
}