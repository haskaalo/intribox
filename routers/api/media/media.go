package media

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/middlewares"
)

// InitRouter with all media endpoint
func InitRouter(r *mux.Router) {
	r.HandleFunc("/new", postNew).Methods("POST")
	r.HandleFunc("/get_mediaurl", getMediaURL).Methods("GET")
	r.Use(middlewares.SetSession, middlewares.RequireSession)
}
