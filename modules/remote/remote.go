package remote

import "io"

// ObjectInfo File infernation output
type ObjectInfo struct {
	Path   string
	Size   int64
	SHA256 string
}

// Remote File storage system
type Remote interface {
	WriteFile(filename string, inFolder string, rdata io.Reader) (*ObjectInfo, error)
	ReadFile(path string) (io.Reader, error)
}
