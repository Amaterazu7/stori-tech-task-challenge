package persistence

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

type S3BucketRepository struct {
	Name     string
	Region   string
	S3Client *s3.S3
}

func NewS3BucketRepository(name string, region string) domain.HandleBucketRepository {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccess := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("AWS_S3_BUCKET_ENDPOINT")
	s3ForcePathStyle := true
	sess, err := session.NewSession(
		&aws.Config{
			Region:           aws.String(region),
			Credentials:      credentials.NewStaticCredentials(accessKey, secretAccess, ""),
			S3ForcePathStyle: &s3ForcePathStyle,
			Endpoint:         &endpoint,
		},
	)
	if err != nil {
		log.Printf("[ERROR] :: Session Problem %v", err)
		// TODO: Perhaps here we should return an error, if we are not enable to connect the session
		// 			return "", errors.New(fmt.Sprintf("Session Problem - %v", err.Error()))
	}

	return &S3BucketRepository{
		Name:     name,
		Region:   region,
		S3Client: s3.New(sess),
	}
}

func (s3Bucket S3BucketRepository) FindFileByName(fileName string) (string, error) {
	rawObject, err := s3Bucket.S3Client.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(s3Bucket.Name),
			Key:    aws.String(fileName),
		},
	)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(rawObject.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to READ fom S3 - %v", err.Error()))
	}

	return buf.String(), nil
}

func (s3Bucket S3BucketRepository) Find() error {
	result, err := s3Bucket.S3Client.ListBuckets(nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to list buckets - %v", err.Error()))
	}

	for _, bucket := range result.Buckets {
		log.Printf(
			" [INFO] :: * [%s] bucket created on [%s]\n",
			aws.StringValue(bucket.Name),
			aws.TimeValue(bucket.CreationDate),
		)
	}

	return nil
}
