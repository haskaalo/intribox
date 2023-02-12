package auth

import (
	"encoding/json"
	"net/http"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/response"
	"github.com/rs/zerolog/log"
)

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	params := new(loginParams)
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		response.InvalidParameter(w, "body")
		return
	}

	user, err := models.LogInUser(params.Email, params.Password)
	if err == models.ErrRecordNotFound {
		response.NotFound(w)
		return
	} else if err != nil {
		log.Warn().Err(err).Msg("Error while trying to login user")
		response.InternalError(w)
		return
	}

	selector, validator, err := models.InitiateSession(user.ID)
	if err != nil {
		log.Warn().Err(err).Msg("Error while trying to initiate session")
		response.InternalError(w)
		return
	}
	apiToken := selector + "." + validator

	response.Respond(w, &response.M{
		"apiToken": apiToken,
	}, http.StatusOK)
}
