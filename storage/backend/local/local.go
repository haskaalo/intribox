package local

import (
	"io"
	"os"
	"path/filepath"

	"github.com/haskaalo/intribox/config"
)

// R exported local backend
type R struct{}

// RemoveFile from local
func (*R) RemoveFile(path string) error {
	return os.Remove(filepath.Join(config.Storage.UserDataPath, path))
}

// ReadFile from local
func (*R) ReadFile(name string) (io.Reader, error) {
	return nil, nil
}
