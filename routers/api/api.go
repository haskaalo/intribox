package api

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/routers/api/album"
	"github.com/haskaalo/intribox/routers/api/auth"
	"github.com/haskaalo/intribox/routers/api/media"
)

// InitRouter Add all API paths to a router
func InitRouter(r *mux.Router) {
	auth.InitRouter(r.PathPrefix("/auth").Subrouter())
	media.InitRouter(r.PathPrefix("/media").Subrouter())
	album.InitRouter(r.PathPrefix("/album").Subrouter())
}
