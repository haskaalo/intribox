package album

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/middlewares"
)

func InitRouter(r *mux.Router) {
	r.HandleFunc("/post_new", postNewAlbum).Methods("POST")
	r.Use(middlewares.SetSession, middlewares.RequireSession)
}
