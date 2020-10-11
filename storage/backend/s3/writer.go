package s3

import (
	"io"

	"github.com/haskaalo/intribox/storage/backend"
)

type s3Writer struct{}

// NewObjectWriter prepare file to be uploaded to s3
// usually to tmp first
func (*R) NewObjectWriter(in io.Reader) (backend.ObjectWriter, error) {
	/*writer := new(s3Writer)

	uploader := s3manager.NewUploader(awsSession)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("something"),
		Key:    aws.String("something to change"),
		Body:   in,
	})*/

	return nil, nil
}
