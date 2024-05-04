package bucket

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

const (
	AwsProvider StorageBucketType = iota
)

type StorageBucketType int

func New(bt StorageBucketType, cfg any) (b *Bucket, err error) {
	rt := reflect.TypeOf(cfg)

	switch bt {
	case AwsProvider:
		// TODO: Implement Aws Provider
	default:
		return nil, fmt.Errorf("unknown storage bucket type: %v", bt)
	}

	return
}

type StorageBucketInterface interface {
	Upload(io.Reader, string) error
	Download(string, string) (*os.File, error)
	Delete(string) error
}

type Bucket struct {
	p StorageBucketInterface
}

func (b *Bucket) Upload(file io.Reader, key string) error {
	return b.p.Upload(file, key)
}

func (b *Bucket) Download(src, dst string) (file *os.File, err error) {
	return b.p.Download(src, dst)
}

func (b *Bucket) Delete(key string) error {
	return b.p.Delete(key)
}
