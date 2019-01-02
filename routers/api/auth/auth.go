package auth

import (
	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/middlewares"
)

// InitRouter Add Auth routes to a Mux Router
func InitRouter(r *mux.Router) {
	r.HandleFunc("/login", postLogin).Methods("POST")
	r.HandleFunc("/logout", postLogout).Methods("POST")

	r.Use(middlewares.SetSession)
}
