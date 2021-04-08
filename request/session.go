package request

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/haskaalo/intribox/models"
)

const privKeySession = 0

// GetSession get session data
func GetSession(r *http.Request) *models.Session {
	rv := context.Get(r, privKeySession)
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
func SetSession(sess *models.Session, r *http.Request) {
	context.Set(r, privKeySession, sess)
}
