package db

import (
	"github.com/sirupsen/logrus"
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
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	t.ID = bson.NewObjectId()
	return ms.DB(DBName).C("topics").Insert(t)
}

/*
ReadAll reads all the topic documents
*/
func (ts *Topics) ReadAll() error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("topics").Find(bson.M{}).Sort("title").All(ts)
}

/*
Update edits a topic document
*/
func (t *Topic) Update(u bson.M) error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}

	_, err = ms.DB(DBName).C("topics").Find(bson.M{"_id": t.ID}).Apply(change, t)
	return err
}

/*
Delete deletes a topic document
*/
func (t *Topic) Delete() error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("topics").Remove(bson.M{"_id": t.ID})
}
