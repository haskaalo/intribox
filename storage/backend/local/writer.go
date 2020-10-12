package local

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/storage/backend"
)

type localWriter struct {
	tmpfile *os.File
	sha256  string
	size    int64
}

// WriteObject prepare file to be uploaded to local
func (*R) WriteObject(in io.Reader, path string) (backend.ObjectAction, error) {
	writer := new(localWriter)
	fullPath := config.Storage.UserDataPath + "/" + path

	err := os.MkdirAll(filepath.Dir(fullPath), 0777)
	if err != nil {
		return nil, err
	}

	osfile, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}

	// Copy song to tempfile
	hasher := sha256.New()
	size, err := io.Copy(osfile, io.TeeReader(in, hasher))
	osfile.Close()
	if err != nil {
		os.Remove(osfile.Name())
		return nil, err
	}

	writer.tmpfile = osfile
	writer.sha256 = hex.EncodeToString(hasher.Sum(nil))
	writer.size = size

	return writer, nil
}

func (w *localWriter) Delete() error {
	return os.Remove(w.tmpfile.Name())
}

func (w localWriter) SHA256() string {
	return w.sha256
}

func (w localWriter) Size() int64 {
	return w.size
}
