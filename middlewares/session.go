package middlewares

import (
	"net/http"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/modules/context"
	"github.com/haskaalo/intribox/response"
)

// SetSession Set Session variable in request
func SetSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get(models.SessionHeaderName)
		if auth == "" {
			next.ServeHTTP(rw, r)
			return
		}

		session, err := models.GetSessionByToken(auth)
		if err == models.ErrNotValidSessionToken {
			next.ServeHTTP(rw, r)
			return
		} else if err != nil {
			response.InternalError(rw)
			return
		}

		context.SetSession(r, session)
		err = session.ResetTimeSession()
		if err != nil {
			response.InternalError(rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
