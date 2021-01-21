package test

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
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
}

// MockServerTearDown Close test server
func MockServerTearDown() {
	defer Server.Close()
}
