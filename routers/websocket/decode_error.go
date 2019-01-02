package websocket

import "reflect"

// IsJSONError tell if error is from package encoding/json
func IsJSONError(err error) bool {
	if tpe := reflect.TypeOf(err).Elem().PkgPath(); tpe == "encoding/json" {
		return true
	}

	return false
}
