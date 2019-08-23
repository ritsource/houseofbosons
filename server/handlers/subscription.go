package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/houseofbosons/houseofbosons/server/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

/*
ReadSubscriptions handles requests to read all the subscription documents from database
*/
func ReadSubscriptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	var ss db.Subscriptions
	err := ss.ReadAll()
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	writeJSON(w, ss)
}

/*
CreateSubscription creates a new subscription document
*/
func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading the json body
	var s db.Subscription
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	if !rxEmail.MatchString(s.Email) {
		writeErr(w, 400, fmt.Errorf("empty email \"%v\"is not valid", s.Email))
		return
	}

	err = s.Read(bson.M{"email": s.Email}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		// ok to go
	case nil:
		writeErr(w, http.StatusConflict, fmt.Errorf("email \"%v\" has already been in use", s.Email))
		return
	default:
		writeErr(w, 500, err)
		return
	}

	// setting `Subscribed` to true and `SubscribedAt` to current  timestamp
	s.Subscribed = true
	s.SubscribedAt = int32(time.Now().Unix())

	// inserting a new document in the database
	err = s.Create()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, s)
}
