package local

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
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

// NewObjectWriter prepare file to be uploaded to local
func (*R) NewObjectWriter(in io.Reader) (backend.ObjectWriter, error) {
	writer := new(localWriter)

	err := os.MkdirAll(config.Storage.UserDataPath+"/tmp", 0777)
	if err != nil {
		return nil, err
	}
	osfile, err := ioutil.TempFile(config.Storage.UserDataPath+"/tmp", "intribox_")
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

func (w *localWriter) Move(path string) error {
	systemPath := filepath.Join(config.Storage.UserDataPath, filepath.Dir(path), filepath.Base(path))
	err := os.MkdirAll(filepath.Dir(systemPath), 0777)
	if err != nil {
		os.Remove(w.tmpfile.Name())
		return err
	}

	err = os.Rename(w.tmpfile.Name(), systemPath)
	if err != nil {
		os.Remove(w.tmpfile.Name())
		return err
	}

	return nil
}

func (w *localWriter) Cancel() {
	os.Remove(w.tmpfile.Name())
}

func (w localWriter) SHA256() string {
	return w.sha256
}

func (w localWriter) Size() int64 {
	return w.size
}
