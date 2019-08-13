package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
Topic struct represents a topic, kind of categories of blogs
*/
type Topic struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string        `bson:"title" json:"title"`
}

/*
Topics is slice of Topic
*/
type Topics []Topic

/*
Create creates a new topic document
*/
func (t *Topic) Create() error {
	t.ID = bson.NewObjectId()
	return MgoDB.C("topics").Insert(t)
}

/*
ReadAll reads all the topic documents
*/
func (ts *Topics) ReadAll() error {
	return MgoDB.C("topics").Find(bson.M{}).Sort("title").All(ts)
}

/*
Update edits a topic document
*/
func (t *Topic) Update(u bson.M) error {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := MgoDB.C("topics").Find(bson.M{"_id": t.ID}).Apply(change, t)

	return err
}

/*
Delete deletes a topic document
*/
func (t *Topic) Delete() error {
	err := MgoDB.C("topics").Remove(bson.M{"_id": t.ID})
	return err
}
