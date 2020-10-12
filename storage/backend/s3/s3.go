package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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

// ReadObject from s3
// TODO
func (*R) ReadObject(name string) (io.Reader, error) {
	return nil, nil
}
