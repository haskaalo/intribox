package response

import (
	"encoding/json"
	"net/http"
)

// M Act as an alias for map[string]interface{}
type M map[string]interface{}

// Respond respond a request with JSON
func Respond(w http.ResponseWriter, in interface{}, status int) {
	jsonData, err := json.Marshal(in)
	if err != nil {
		InternalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
