package remote

import "io"

// Remote File storage system
type Remote interface {
	// NewObjectWriter prepare file to be uploaded to a remote
	NewObjectWriter(in io.Reader) (ObjectWriter, error)

	RemoveFile(path string) error

	ReadFile(path string) (io.Reader, error)
}

// ObjectWriter Upload prepared (tmp) file to remote
type ObjectWriter interface {
	ObjectInfo

	// Move file from tmp to local directory
	Move(path string) error

	// Cancel file write by removing tmp file
	Cancel()
}

// ObjectInfo object related info
type ObjectInfo interface {
	SHA256() string

	Size() int64
}
