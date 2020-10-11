package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSession *session.Session

// R exported s3 backend
type R struct{}
