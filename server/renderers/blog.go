package renderers

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/houseofbosons/houseofbosons/server/db"
	"gopkg.in/mgo.v2/bson"
)

/*
PageIdxes holds values related to page-navigation in `/posts` handler.
Last, is the lastest possible page, and Current represents the currently
requested page for by the client
*/
type PageIdxes struct {
	Last    int
	Current int
}

/*
BlogsHandler renders lists out all the blogs as UI for the `/posts`
page, renders only the public `blogs` and includes page navigation
*/
func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	nbpp := 3          // max number of blog-posts to be shown in 1 page
	num := PageIdxes{} // num holds page index related daa

	// reading the page-number from query staring
	var err error
	num.Current, _ = strconv.Atoi(r.URL.Query().Get("pagenum"))
	topic := r.URL.Query().Get("topic")

	// bs represents the slice of blogs
	var bs db.Blogs

	// reading the number of documents from database
	nbs, err := bs.Count(bson.M{"is_deleted": false, "is_public": true})
	if err != nil {
		renderErr(w, 422, "Unable to Read from Database")
		return
	}

	// the last possible page-index
	num.Last = int(math.Ceil(float64(nbs / nbpp)))

	// if requested page-index doesn't exist the set it back to 0
	if num.Current > num.Last {
		num.Current = num.Last
	}

	// reading all the requested blogs posts data (only the ones taht exists in this page)
	err = bs.ReadFew(bson.M{"is_deleted": false, "is_public": true}, bson.M{}, num.Current*nbpp, nbpp)
	if err != nil {
		renderErr(w, 422, "Unable to Read Blogs from Database")
		return
	}

	// reading `topic` documents from the database
	var ts db.Topics
	err = ts.ReadAll()
	if err != nil {
		renderErr(w, 422, "Unable to Read Topics from Database")
		return
	}

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/posts.html",
		"static/partials/posts-item.html",
		"static/partials/header.html",
		"static/partials/footer.html",
	)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	err = t.Execute(w, struct {
		// NavIdxes [][]int
		NavArray   []NavArray
		NavCurrent int
		Posts      db.Blogs
		Topics     db.Topics
		TopicSel   string
	}{
		NavArray:   getNavArray(num),
		NavCurrent: num.Current,
		Posts:      bs,
		Topics:     ts,
		TopicSel:   topic,
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}

/*
NavArray struct holds each element
of nav-tool to be shown in the view
*/
type NavArray struct {
	B bool
	V int
}

/*
getNavArray returns the desired navigation page number array
*/
func getNavArray(idxs PageIdxes) []NavArray {
	arr := []NavArray{}

	var lf int8 // int that keep tracks how many of last iretations are consecutively false
	var cf bool // bool that if more false to be added in the array or not

	for i := 0; i < idxs.Last+1; i++ {
		if i-1 == idxs.Current || i == idxs.Current || i+1 == idxs.Current {
			arr = append(arr, NavArray{B: true, V: i})
			lf = 0
			cf = true
		} else if i+1 == idxs.Last || i == idxs.Last {
			arr = append(arr, NavArray{B: true, V: i})
			lf = 0
			cf = true
		} else if i == 0 {
			arr = append(arr, NavArray{B: true, V: i})
			lf = 0
			cf = true
		} else {
			if cf {
				arr = append(arr, NavArray{B: false, V: i})
			}

			lf++
		}

		if lf == 1 {
			cf = false
		}
	}

	return arr
}
