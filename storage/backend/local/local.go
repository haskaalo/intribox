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
func (*R) readObject(path string) (io.Reader, error) {
	fullPath := config.Storage.UserDataPath + "/" + path

	file, err := os.OpenFile(filepath.Dir(fullPath), os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// GetReadObjectURL Return an path similar to models.GetSongPath (e.g.: localhost:8080/api/storage/2/song/61230e8e-896d-4380-b00e-64364e79cad5)
func (*R) GetReadObjectURL(path string) (string, error) {
	return (config.Server.Hostname + "/api/storage" + path), nil
}
