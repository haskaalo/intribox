package album

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/middlewares"
)

func InitRouter(r *mux.Router) {
	r.HandleFunc("/new", postNewAlbum).Methods("POST")
	r.HandleFunc("/list", getAlbumList).Methods("GET")

	r.Use(middlewares.SetSession, middlewares.RequireSession)
}
