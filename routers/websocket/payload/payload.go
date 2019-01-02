package payload

import (
	"encoding/json"
)

// Base for sending and receiving messages
type Base struct {
	Data json.RawMessage `json:"d"`
	Type string          `json:"t"`
}

// UnmarshalBase parse Base payload
func UnmarshalBase(data []byte) (base Base, err error) {
	err = json.Unmarshal(data, base)
	if err != nil {
		return
	}
	return base, err
}
