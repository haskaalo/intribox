package backend

import "io"

// Backend interface used by all storage system
type Backend interface {
	// NewObjectWriter prepare file to be uploaded to a storage system
	NewObjectWriter(in io.Reader) (ObjectWriter, error)

	RemoveFile(path string) error

	ReadFile(path string) (io.Reader, error)
}

// ObjectWriter Upload prepared (tmp) file to storage
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
