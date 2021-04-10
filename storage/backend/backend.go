package backend

import "io"

// Backend interface used by all storage system
type Backend interface {
	// WriteObject prepare file to be uploaded to a storage system
	WriteObject(in io.Reader, path string) (ObjectAction, error)

	RemoveObject(path string) error

	GetReadObjectURL(path string, MediaID int) (string, error)

	ReadObject(path string) (io.Reader, error)
}

// ObjectAction Upload prepared (tmp) file to storage
type ObjectAction interface {
	ObjectInfo

	// Cancel file write by removing tmp file
	Delete() error
}

// ObjectInfo object related info
type ObjectInfo interface {
	SHA256() string

	Size() int64
}
