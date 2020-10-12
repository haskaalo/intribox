package local

import (
	"io"
	"os"
	"path/filepath"

	"github.com/haskaalo/intribox/config"
)

// R exported local backend
type R struct{}

// RemoveObject from local
func (*R) RemoveObject(path string) error {
	return os.Remove(filepath.Join(config.Storage.UserDataPath, path))
}

// ReadObject from local
func (*R) ReadObject(name string) (io.Reader, error) {
	return nil, nil
}
