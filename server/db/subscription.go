package db

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
Subscription struct represents a subscription (simply email for newsletter)
*/
type Subscription struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Email        string        `bson:"email" json:"email"`
	Subscribed   bool          `bson:"subscribed" json:"subscribed"`
	SubscribedAt int32         `bson:"subscribed_at" json:"subscribed_at"`
}

/*
Subscriptions is slice of `Subscription`
*/
type Subscriptions []Subscription

/*
Create creates a new subscriptions document
*/
func (s *Subscription) Create() error {
	ms, err := DBConnect()
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	s.ID = bson.NewObjectId()
	return ms.DB(DBName).C("subscriptions").Insert(s)
}

/*
ReadAll reads all the subscription documents
*/
func (ss *Subscriptions) ReadAll() error {
	ms, err := DBConnect()
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("subscriptions").Find(bson.M{}).Sort("-created_at").All(ss)
}

// Read - Reads single subscription document
func (s *Subscription) Read(f, sel bson.M) error {
	ms, err := DBConnect()
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("subscriptions").Find(f).Select(sel).One(s)
}

/*
Update edits a subscription document
*/
func (s *Subscription) Update(u bson.M) error {
	ms, err := DBConnect()
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}

	_, err = ms.DB(DBName).C("subscriptions").Find(bson.M{"_id": s.ID}).Apply(change, s)
	return err
}

/*
Delete deletes a subscription document
*/
func (s *Subscription) Delete() error {
	ms, err := DBConnect()
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("subscriptions").Remove(bson.M{"_id": s.ID})
}
