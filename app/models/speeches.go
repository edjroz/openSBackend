package models

import "gopkg.in/mgo.v2/bson"

/*
Speeches model
*/
type Speeches struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	
}