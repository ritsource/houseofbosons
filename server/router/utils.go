package router

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func writeErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)                                     // Status Code
	w.Header().Set("Content-Type", "application/json")        // Response Type - JSON
	w.Write([]byte("{\"message\": \"" + err.Error() + "\"}")) // Error Message

	logrus.Errorf("%v\n", err)
}
