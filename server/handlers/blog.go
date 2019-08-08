/*
`blog.go` contains handler function for houseofboson-blogs-api
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/houseofbosons/houseofbosons/server/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
CreateBlog creates a new blog from provided data read on `http.Request` body,
the body needs to contain all required fields in json, otherwise it will be
saved as `null` in the database
*/
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading the json body, provided in the request
	var b db.Blog
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// manually setting `CreatedAt`, that holds creation time
	b.CreatedAt = int32(time.Now().Unix())

	// inserting a new document in the database
	err = b.Create()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// redirecting to `/api/private/blogs` route handler
	http.Redirect(w, r, "/api/private/blogs", http.StatusTemporaryRedirect) // 302 - POST to GET
}

/*
ReadBlog reads a a single blog from the database by
it's `ID` and writes the back the data in json
*/
func ReadBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from `query-string`
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// reading the document from database
	var b db.Blog
	err := b.Read(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		writeErr(w, 404, err)
	case nil:
		// everything's fine
	default:
		writeErr(w, 500, err)
		return
	}

	// writing json data to the client
	writeJSON(w, b)
}

/*
ReadBlogs reads all the `blogs` documents from database without any filter,
and that's why only an authenticated admin client is allowed to access that
If `skip` and `limit` value provided in the query string, it only fetches
those few blog-documents. If not provided then fetches all blogs documents
*/
func ReadBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	/*
		reading `skip` and `limit` values from the query string.
		If none provided then it fetches all the blog documents,
		else it fetches only the documents between `index-skip`
		and `index-(skip+limit)`. If throws error if invalid
		value provided in the query (non-numeric)
	*/
	skpstr := r.URL.Query().Get("skip")
	limstr := r.URL.Query().Get("limit")

	var bs db.Blogs

	/*
		checking if both `skip` and `limit` has been provided, if not
		then read all blog documents, else reading only the required
	*/
	if len(skpstr) == 0 || len(limstr) == 0 {
		err := bs.ReadAll(bson.M{}, bson.M{})
		if err != nil {
			writeErr(w, 500, err)
			return
		}
	} else {
		// integer value of query
		skp, err := strconv.Atoi(skpstr)
		lim, err := strconv.Atoi(limstr)

		if err != nil {
			writeErr(w, 400, fmt.Errorf("invalid skip and limit value provided"))
			return
		}

		// reading the required documents
		err = bs.ReadFew(bson.M{}, bson.M{}, skp, lim)
		if err != nil {
			writeErr(w, 500, err)
			return
		}
	}

	// writing teh data back to the client
	writeJSON(w, bs)
}

/*
EditBlog edits a blog document
*/
func EditBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// reading request body
	var body map[string]interface{} // because cannot use b (type db.Blog) as type bson.M in argument to bl.Update
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// editing document
	var b db.Blog
	err = b.Update(bson.M{"_id": bson.ObjectIdHex(id)}, body) // Update Document in Database
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, b) // Write Data
}

/*
DeleteBlog deleted a blog (not permanently) by updating
the document's `is_deleted` property to `true`
*/
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// editing document to `is_deleted: true`
	var b db.Blog
	err := b.Delete(bson.ObjectIdHex(id))
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// writing the updated data
	writeJSON(w, b)
}

/*
DeleteBlogPrem permanently deletes a blog document from the database
*/
func DeleteBlogPrem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// deleting document (permanently)
	var b db.Blog
	b.ID = bson.ObjectIdHex(id)
	err := b.DeletePermanent()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}

/*
CheckIDStr .
*/
func CheckIDStr(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	idstr := r.URL.Query().Get("idstr")
	if len(idstr) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no idstr provided"))
		return
	}

	// the response
	var res struct {
		Av bool `json:"available"`
	}

	// deleting document (permanently)
	var b db.Blog
	err := b.Read(bson.M{"id_str": idstr}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		res.Av = true
	case nil:
		res.Av = false
	default:
		writeErr(w, 500, err)
		return
	}

	// writing the updated data
	writeJSON(w, res)
}
