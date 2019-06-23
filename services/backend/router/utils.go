package router

import "net/http"

func writeError(w http.ResponseWriter, status int, err error, msg string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\",\"error\": \"" + err.Error() + "\"}"))
}
