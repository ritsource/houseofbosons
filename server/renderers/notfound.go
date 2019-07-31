package renderers

import (
	"fmt"
	"net/http"
)

// NotFoundHandler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "page not found")
}
