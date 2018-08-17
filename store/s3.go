package store

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 Store definition
type S3 struct {
	client *s3.S3
	Bucket string
	Prefix string
}

// NewS3Store build a new S3 store
func NewS3Store(httpClient *http.Client) (Store, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	session, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			HTTPClient: httpClient,
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return S3{}, err
	}
	return S3{
		client: s3.New(session),
		Prefix: os.Getenv("slssl_S3_PREFIX"),
		Bucket: os.Getenv("slssl_S3_BUCKET"),
	}, nil
}

// FetchFile a file from s3
func (store S3) FetchFile(filename string, buf *bytes.Buffer) error {
	s3Key := store.Prefix + filename

	s3Object, err := store.client.GetObject(&s3.GetObjectInput{
		Bucket: &store.Bucket,
		Key:    &s3Key,
	})
	if err != nil {
		return err
	}

	defer s3Object.Body.Close()
	if _, err := io.Copy(buf, s3Object.Body); err != nil {
		return err
	}

	return nil
}

// PutFile put a file into the S3 store
func (store S3) PutFile(filename string, fileReader *bytes.Reader) error {
	s3Key := store.Prefix + filename

	_, err := store.client.PutObject(&s3.PutObjectInput{
		Bucket: &store.Bucket,
		Key:    &s3Key,
		Body:   fileReader,
	})

	return err
}
