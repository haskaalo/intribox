package s3

import (
	"errors"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/haskaalo/intribox/config"
)

var s3Client *s3.S3

// R exported s3 backend
type R struct{}

func init() {
	s3Client = s3.New(config.AwsSession)
}

// RemoveObject from S3 bucket
func (*R) RemoveObject(path string) error {
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.Aws.Bucket),
		Key:    aws.String(path),
	})

	return err
}

// GetReadObjectURL return a presigned URL so users can download file directly from S3 without reaching our server, thus resulting less bandwidth usage
func (*R) GetReadObjectURL(path string, MediaID uuid.UUID) (string, error) {
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.Aws.Bucket),
		Key:    aws.String(path),
	})
	url, err := req.Presign(5 * time.Minute)

	return url, err
}

// ReadObject should not be called when the remote is S3
func (*R) ReadObject(path string) (io.Reader, error) {
	return nil, errors.New("ReadObject should not be called when the remote is S3")
}
