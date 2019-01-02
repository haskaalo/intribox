package middlewares

import (
	"net/http"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/context"
	"github.com/haskaalo/intribox/response"
)

// RequireSession return unauthorized in JSON if user is not logged in
func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session := context.GetSession(r)
		if (models.Session{}) == *session {
			response.Unauthorized(rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
