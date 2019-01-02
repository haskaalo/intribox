package auth

import (
	"net/http"

	"github.com/haskaalo/intribox/response"
)

func getSignIn(w http.ResponseWriter, r *http.Request) {
	response.Render(w, response.RenderData{
		Title: "IntroBox - Sign In",
	})
}
