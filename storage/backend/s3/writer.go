package s3

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/storage/backend"
)

type s3Writer struct {
	key    string
	sha256 string
	size   int64
}

// WriteObject prepare file to be uploaded to s3
// usually to tmp first
func (*R) WriteObject(in io.Reader, path string) (backend.ObjectAction, error) {
	s3writer := new(s3Writer)

	hasher := sha256.New() // To get the SHA256 hash of the file

	// Upload file to S3 bucket
	uploader := s3manager.NewUploader(config.AwsSession)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Aws.Bucket),
		Key:    aws.String(path),
		Body:   io.TeeReader(in, hasher),
	})
	if err != nil {
		return nil, err
	}

	// Get size of the file by doing a HEAD request
	bodyResult, err := s3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(config.Aws.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		var prevErr error
		_, prevErr = s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(config.Aws.Bucket),
			Key:    aws.String(path),
		})
		if prevErr != nil { // To avoid override previous error with nil if deleteobject is successful
			return nil, prevErr
		}
		return nil, err
	}

	s3writer.key = path
	s3writer.sha256 = hex.EncodeToString(hasher.Sum(nil))
	s3writer.size = *bodyResult.ContentLength

	return s3writer, nil
}

func (w *s3Writer) Delete() error {
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.Aws.Bucket),
		Key:    aws.String(w.key),
	})
	if err != nil {
		return err
	}

	return nil
}

func (w s3Writer) SHA256() string {
	return w.sha256
}

func (w s3Writer) Size() int64 {
	return w.size
}
