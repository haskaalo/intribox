package local

import (
	"io"
	"os"
	"path/filepath"

	"github.com/haskaalo/intribox/config"
)

// R local
type R struct{}

// RemoveFile from local
func (r R) RemoveFile(path string) error {
	return os.Remove(filepath.Join(config.Storage.UserDataPath, path))
}

// ReadFile from local
func (r R) ReadFile(name string) (io.Reader, error) {
	return nil, nil
}
