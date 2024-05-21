package bucket

import (
	"fmt"
	"io"
	"reflect"
)

const (
	AwsProvider StorageBucketType = iota
	MockProvider
)

type StorageBucketType int

func New(bt StorageBucketType, cfg any) (b *Bucket, err error) {
	b = new(Bucket)
	rt := reflect.TypeOf(cfg)

	switch bt {
	case AwsProvider:
		if rt.Name() != "AwsConfig" {
			return nil, fmt.Errorf("configuration must be of type AwsConfig")
		}

		b.p = newAwsSession(cfg.(AwsConfig))
	case MockProvider:
		b.p = &MockBucket{
			content: make(map[string][]byte),
		}
	default:
		return nil, fmt.Errorf("unknown storage bucket type: %v", bt)
	}

	return
}

type StorageBucketInterface interface {
	Upload(io.Reader, string) error
	Download(string, string) error
	Delete(string) error
}

type Bucket struct {
	p StorageBucketInterface
}

func (b *Bucket) Upload(file io.Reader, key string) error {
	return b.p.Upload(file, key)
}

func (b *Bucket) Download(src, dst string) error {
	return b.p.Download(src, dst)
}

func (b *Bucket) Delete(key string) error {
	return b.p.Delete(key)
}
