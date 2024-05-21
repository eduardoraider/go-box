package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/eduardoraider/go-box/internal/bucket"
	"github.com/eduardoraider/go-box/internal/queue"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// rabbitmq config
	qcfg := queue.RabbitMQConfig{
		URL:       "amqp://" + os.Getenv("RABBITMQ_URL"),
		TopicName: os.Getenv("RABBITMQ_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	// create new queue
	qc, err := queue.New(queue.RabbitMQ, qcfg)
	if err != nil {
		panic(err)
	}

	// create channel to consume messages
	c := make(chan queue.AppQueueDto)
	go qc.Consume(c)

	// bucket config
	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				""),
		},
		BucketDownload: "wookye-gobox-raw",
		BucketUpload:   "wookye-gobox-gzip",
	}

	// create new bucket session
	b, err := bucket.New(bucket.AwsProvider, bcfg)
	if err != nil {
		panic(err)
	}

	log.Printf("Waiting for messages")

	for msg := range c {

		dst := fmt.Sprintf("%d_%s", msg.ID, msg.Filename)

		log.Printf("Start working on %s", msg.Filename)

		err := b.Download(msg.Path, dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		file, err := os.Open(dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		body, err := io.ReadAll(file)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, err = zw.Write(body)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		if err := zw.Close(); err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		err = b.Upload(zr, msg.Path)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		err = os.Remove(dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		log.Printf("%s was proccessed successfully", msg.Filename)
	}
}
