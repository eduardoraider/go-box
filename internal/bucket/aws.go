package bucket

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"io"
	"os"
)

type AwsConfig struct {
	Config         *aws.Config
	BucketDownload string
	BucketUpload   string
}

func newAwsSession(cfg AwsConfig) *awsSession {
	c := session.New(cfg.Config)

	return &awsSession{
		sess:           c,
		bucketDownload: cfg.BucketDownload,
		bucketUpload:   cfg.BucketUpload,
	}
}

type awsSession struct {
	sess           *session.Session
	bucketDownload string
	bucketUpload   string
}

func (as *awsSession) Download(src, dst string) (file *os.File, err error) {
	return nil, err
}

func (as *awsSession) Upload(file io.Reader, key string) (err error) {
	return nil
}

func (as *awsSession) Delete(key string) error {
	return nil
}
