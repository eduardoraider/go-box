package bucket

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func (as *awsSession) Download(src, dst string) error {
	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(as.sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(as.bucketDownload),
			Key:    aws.String(src),
		})

	return nil
}

func (as *awsSession) Upload(file io.Reader, key string) (err error) {
	uploader := s3manager.NewUploader(as.sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(as.bucketUpload),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func (as *awsSession) Delete(key string) error {
	svc := s3.New(as.sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(as.bucketDownload),
		Key:    aws.String(key),
	})

	if err != nil {
		return err
	}

	return svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(as.bucketDownload),
		Key:    aws.String(key),
	})
}
