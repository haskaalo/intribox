package local

import (
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/haskaalo/intribox/config"
)

// R exported local backend
type R struct{}

// RemoveObject from local
func (*R) RemoveObject(path string) error {
	return os.Remove(filepath.Join(config.Storage.UserDataPath, path))
}

// ReadObject from local
func (*R) ReadObject(path string) (io.Reader, error) {
	fullPath := config.Storage.UserDataPath + "/" + path

	file, err := os.OpenFile(filepath.Dir(fullPath), os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// GetReadObjectURL Return an path similar to models.GetMediaPath (e.g.: localhost:8080/api/media/download?mediaid=34567855)
func (*R) GetReadObjectURL(path string, mediaID int) (string, error) {
	return (config.Server.Hostname + "/api/media/download?mediaid=" + strconv.Itoa(mediaID)), nil
}
