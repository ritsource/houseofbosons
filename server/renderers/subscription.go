package renderers

import (
	"net/http"
	"text/template"
)

// SubscribeHandler .
func SubscribeHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles(
		"static/pages/subscribe.html",
		"static/partials/header.html",
		"static/partials/footer.html",
		"static/partials/head-links.html",
	)
	if err != nil {
		writeErr(w, 500, err)
	}

	err = t.Execute(w, []string{})
	if err != nil {
		writeErr(w, 500, err)
	}
}
