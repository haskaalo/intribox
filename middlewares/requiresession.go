package middlewares

import (
	"net/http"

	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
)

// RequireSession return unauthorized in JSON if user is not logged in
func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session := request.GetSession(r)
		if session == nil {
			response.Unauthorized(rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
