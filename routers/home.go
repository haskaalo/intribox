package routers

import (
	"net/http"

	"github.com/haskaalo/intribox/response"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	response.Render(w, response.RenderData{
		Title: "intribox - Home",
	})
}
