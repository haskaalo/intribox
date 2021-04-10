package test

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/haskaalo/intribox/config"
)

var (
	// Router is the HTTP request multiplexer used with the test server.
	Router *mux.Router

	// Server is a test HTTP server used to provide mock API responses.
	Server *httptest.Server
)

// MockServerSetup Setup test server
func MockServerSetup() {
	Router = mux.NewRouter()
	Server = httptest.NewServer(Router)
	config.Server.Hostname = Server.URL
}

// MockServerTearDown Close test server
func MockServerTearDown() {
	defer Server.Close()
}
