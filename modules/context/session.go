package context

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/haskaalo/intribox/models"
)

const privKeySession = 0

// GetSession get session data
func GetSession(r *http.Request) *models.Session {
	sess := &models.Session{}
	if rv := context.Get(r, privKeySession); rv != nil {
		sess = rv.(*models.Session)
	}
	return sess
}

// SetSession Set session data on variable
func SetSession(r *http.Request, sess *models.Session) {
	context.Set(r, privKeySession, sess)
}
