package persistence

import (
	"fmt"
	"github.com/amaterazu7/transaction-processor/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type S3Client struct {
	Region string
	Sess   *session.Session
	Svc    *s3.S3
}

type S3BucketRepository struct {
	Region   string
	S3Client *s3.S3
}

func NewS3BucketRepository(region string) domain.HandleBucketRepository {
	// endpoint := os.Getenv("AWS_ENDPOINT")
	// endpoint := os.Getenv("AWS_ACCESS_KEY_ID")
	// endpoint := os.Getenv("AWS_SECRET_ACCESS_KEY")
	// endpoint := "http://localhost:4569"
	endpoint := "http://host.docker.internal:4566"
	s3ForcePathStyle := true
	sess, err := session.NewSession(
		&aws.Config{
			Region:           aws.String(region),
			Credentials:      credentials.NewStaticCredentials("S3RVER", "S3RVER", ""),
			S3ForcePathStyle: &s3ForcePathStyle,
			Endpoint:         &endpoint,
		},
	)
	if err != nil {
		log.Printf("[ERROR] :: Session Problem %v", err)
	}

	return &S3BucketRepository{
		Region:   region,
		S3Client: s3.New(sess),
	}
}

func (s3Bucket S3BucketRepository) FindFileByName(fileName string) error {

	log.Printf("[INFO] :: NAME :: %s", fileName)
	result, err := s3Bucket.S3Client.ListBuckets(nil) // GetObject()
	if err != nil {
		log.Printf("[ERROR] :: Unable to list buckets %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	return nil
}
