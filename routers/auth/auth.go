package auth

import "github.com/gorilla/mux"

// InitRouter add Auth routes to a Mux router
func InitRouter(r *mux.Router) {
	r.HandleFunc("/sign_in", getSignIn).Methods("GET")
}
