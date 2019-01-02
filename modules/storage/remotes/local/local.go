package local

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/haskaalo/intribox/modules/storage/remotes"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/modules/bodyanalyzer"
)

// R local
type R struct{}

// WriteFile to local and return useful data
// inFolder can be for example: /{:userid}/songs
func (R) WriteFile(filename string, inFolder string, rdata io.Reader) (*remotes.ObjectInfo, error) {
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
	analyzer := bodyanalyzer.New(sha256.New(), true)
	_, err = io.Copy(osfile, io.TeeReader(rdata, analyzer))
	osfile.Close()
	if err != nil {
		os.Remove(osfile.Name())
		return nil, err
	}

	// Move the file from tmp to the actual dest with the actual name
	fileHash := analyzer.HexHash()

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

	return &remotes.ObjectInfo{
		Path:   filepath.ToSlash(pathDirectory),
		Size:   analyzer.Size(),
		SHA256: fileHash,
	}, nil
}

// ReadFile from local
func (R) ReadFile(name string) (io.Reader, error) {
	return nil, nil
}

func parseName(name string) (directory string, filename string) {
	subString := strings.Split(name, "/")

	for idx, dirpath := range subString {
		if idx == len(subString)-1 {
			filename = dirpath
			break
		}

		directory += "/" + dirpath
	}

	return
}
