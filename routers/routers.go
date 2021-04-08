package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/response"
	"github.com/haskaalo/intribox/routers/api"
	"github.com/haskaalo/intribox/routers/auth"
)

// BootstrapRouters Create a Router with every intribox paths
func BootstrapRouters() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(config.Client.AssetsPath+"/static"))))

	auth.InitRouter(router.PathPrefix("/auth").Subrouter())
	router.HandleFunc("/home", getHome)

	api.InitRouter(router.PathPrefix("/api").Subrouter())

	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	response.NotFound(w)
}
