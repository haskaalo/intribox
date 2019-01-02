package auth

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/haskaalo/intribox/models"
	"github.com/haskaalo/intribox/response"
)

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.InternalError(w)
		return
	}

	params := new(loginParams)
	err = json.Unmarshal(body, params)
	if err != nil {
		// Probably change depending on the error.
		response.InternalError(w)
		return
	}
	user, err := models.GetUserByEmail(params.Email)

	if err == sql.ErrNoRows {
		response.NotFound(w)
		return
	} else if err != nil {
		response.InternalError(w)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(params.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		response.NotFound(w)
		return
	} else if err != nil {
		response.InternalError(w)
		return
	}

	selector, validator, err := models.InitiateSession(user.ID)
	if err != nil {
		response.InternalError(w)
		return
	}
	apiToken := selector + "." + validator

	response.Respond(w, &response.M{
		"apiToken": apiToken,
	}, http.StatusOK)
}
