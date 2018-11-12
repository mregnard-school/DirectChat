package app

import (
	"net/http"
	u "server/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u.Respond(w, u.Message(false, "This resources was not found on our server", http.StatusNotFound))
		next.ServeHTTP(w, r)
	})
}
