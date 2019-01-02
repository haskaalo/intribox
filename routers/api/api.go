package api

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/routers/api/auth"
	"github.com/haskaalo/intribox/routers/api/song"
)

// InitRouter Add all API paths to a router
func InitRouter(r *mux.Router) {
	auth.InitRouter(r.PathPrefix("/auth").Subrouter())
	song.InitRouter(r.PathPrefix("/song").Subrouter())
}
