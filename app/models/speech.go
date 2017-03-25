package models

import (
	"openss/api/app/models/mongodb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Speech struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Author      string        `json:"author" bson:"author"`
	Time        string        `json:"time" bson:"time"`
	GithubAcc   string        `json:"github_acc" bson:"github_acc"`
	TwitterAcc  string        `json:"twitter_acc" bson:"twitter_acc"`
	Status      bool          `json:"status" bson:"status"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

func newSpeechCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("speeches")
}

// AddSpeech insert a new Speech into database and returns
// last inserted speech on success.
func AddSpeech(m Speech) (speech Speech, err error) {
	c := newSpeechCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateSpeech update a Speech into database and returns
// last nil on success.
func (m Speech) UpdateSpeech() error {
	c := newSpeechCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"title": m.Title, "description": m.Description, "author": m.Author, "time": m.Time, "github_acc": m.GithubAcc, "twitter_acc": m.TwitterAcc, "status": m.Status, "updatedAt": time.Now()},
	})
	return err
}

// DeleteSpeech Delete Speech from database and returns
// last nil on success.
func (m Speech) DeleteSpeech() error {
	c := newSpeechCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetSpeeches Get all Speech from database and returns
// list of Speech on success
func GetSpeeches() ([]Speech, error) {
	var (
		speeches []Speech
		err      error
	)

	c := newSpeechCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&speeches)
	return speeches, err
}

// GetSpeech Get a Speech from database and returns
// a Speech on success
func GetSpeech(id bson.ObjectId) (Speech, error) {
	var (
		speech Speech
		err    error
	)

	c := newSpeechCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&speech)
	return speech, err
}
