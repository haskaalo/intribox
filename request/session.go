package request

import (
	"context"
	"net/http"

	"github.com/haskaalo/intribox/models"
)

type contextKey string

const sessionKey = contextKey("session")

// GetSession get session data
func GetSession(r *http.Request) *models.Session {
	rv := r.Context().Value(sessionKey)
	if rv == nil {
		return nil
	}

	sess, ok := rv.(*models.Session)
	if !ok {
		return nil
	}

	return sess
}

// SetSession Set session data on variable
func SetSession(sess *models.Session, r *http.Request) *context.Context {
	ctx := context.WithValue(r.Context(), sessionKey, sess)
	return &ctx
}
