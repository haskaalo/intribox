package song

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/middlewares"
)

// InitRouter with all song endpoint
func InitRouter(r *mux.Router) {
	r.HandleFunc("/new", postNew).Methods("POST")
	r.HandleFunc("/get_songurl", getSongURL).Methods("GET")
	r.Use(middlewares.SetSession, middlewares.RequireSession)
}
