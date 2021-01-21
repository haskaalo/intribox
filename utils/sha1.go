package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 Hash data into SHA1 hex
func SHA1(n []byte) string {
	h := sha1.New()
	h.Write(n)

	return hex.EncodeToString(h.Sum(nil))
}
