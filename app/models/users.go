package models

import "gopkg.in/mgo.v2/bson"

/*
Users model
*/
type Users struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	
}