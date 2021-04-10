package auth

import (
	"net/http"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/request"
	"github.com/haskaalo/intribox/response"
)

func postLogout(w http.ResponseWriter, r *http.Request) {
	session := request.GetSession(r)
	if session == nil {
		response.Unauthorized(w)
		return
	}

	err := models.DeleteSessionBySelector(session.Selector)

	if err != nil {
		response.InternalError(w)
		return
	}

	response.Respond(w, &response.M{
		"ok": 1,
	}, http.StatusOK)
}
