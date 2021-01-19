package request

import "net/http"

// RequireContentType return bool if expected content-type is set on request
func RequireContentType(expectCt []string, r *http.Request) bool {
	reqContentType := r.Header.Get("Content-Type")
	for _, header := range expectCt {
		if reqContentType == header {
			return true
		}
	}
	return false
}
