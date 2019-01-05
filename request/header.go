package request

import "net/http"

// RequireContentType return bool if expected content-type is set on request
func RequireContentType(expectCt string, r *http.Request) bool {
	return expectHeader("Content-Type", expectCt, r)
}

func expectHeader(headerName string, expected string, r *http.Request) bool {
	header := r.Header.Get(headerName)
	if header != expected {
		return false
	}

	return true
}
