package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ritsource/houseofbosons/server/db"
	"gopkg.in/mgo.v2/bson"
)

/*
ReadTopics handles requests to read all teh topics from database
*/
func ReadTopics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	var ts db.Topics
	err := ts.ReadAll()
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	writeJSON(w, ts)

}

/*
CreateTopic creates a new topic
*/
func CreateTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading the json body
	var t db.Topic
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	if t.Title == "" {
		writeErr(w, 400, fmt.Errorf("empty title \"%v\"is not valid", t.Title))
		return
	}

	// inserting a new document in the database
	err = t.Create()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, t)
}

/*
EditTopic handler edits a topic
*/
func EditTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no topic-id provided"))
		return
	}

	// reading request body
	var body map[string]interface{} // because cannot use t (type db.Topic) as type bson.M in argument to t.Update
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// editing document
	var t db.Topic
	t.ID = bson.ObjectIdHex(id)
	err = t.Update(body)
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, t)
}

/*
DeleteTopic deletes a topic document
*/
func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no topic-id provided"))
		return
	}

	var t db.Topic
	t.ID = bson.ObjectIdHex(id)
	err := t.Delete()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
