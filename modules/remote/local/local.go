package local

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/modules/remote"
)

// R local
type R struct{}

// WriteFile to local and return useful data
// inFolder can be for example: /{:userid}/songs
func (r R) WriteFile(filename string, inFolder string, rdata io.Reader) (*remote.ObjectInfo, error) {
	err := os.MkdirAll(config.Storage.UserDataPath+"/tmp", 0777)
	if err != nil {
		return nil, err
	}

	// Tempfile is created on ${prefixedDir}/tmp in case the system default tmp folder is on another drive
	// If it was on another drive it would cause a actual copy, causing a memory usage rise
	osfile, err := ioutil.TempFile(config.Storage.UserDataPath+"/tmp", "song_")
	if err != nil {
		return nil, err
	}

	// Copy song to tempfile
	hasher := sha256.New()
	size, err := io.Copy(osfile, io.TeeReader(rdata, hasher))
	osfile.Close()
	if err != nil {
		os.Remove(osfile.Name())
		return nil, err
	}

	// Move the file from tmp to the actual dest with the actual name
	fileHash := hex.EncodeToString(hasher.Sum(nil))

	pathDirectory := filepath.Join(inFolder, fileHash[0:2], fileHash[2:4], fileHash+filepath.Ext(filename)) // Path to return in ObjectInfo
	systemPathDirectory := filepath.Join(config.Storage.UserDataPath, pathDirectory)

	err = os.MkdirAll(filepath.Dir(systemPathDirectory), 0777)
	if err != nil {
		os.Remove(osfile.Name())
		return nil, err
	}

	err = os.Rename(osfile.Name(), systemPathDirectory)
	if err != nil {
		os.Remove(osfile.Name())
		return nil, err
	}

	return &remote.ObjectInfo{
		Path:   filepath.ToSlash(pathDirectory),
		Size:   size,
		SHA256: fileHash,
	}, nil
}

// RemoveFile from local
func (r R) RemoveFile(path string) error {
	return os.Remove(filepath.Join(config.Storage.UserDataPath, path))
}

// ReadFile from local
func (r R) ReadFile(name string) (io.Reader, error) {
	return nil, nil
}
