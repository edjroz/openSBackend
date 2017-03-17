package models

import "gopkg.in/mgo.v2/bson"

/*
Speeches Request model
*/
type Rspeech struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	
}