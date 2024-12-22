package utils

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	*s3.Client
}

// deal with path style
func GetS3Client() *S3Client {
	s3Client := s3.NewFromConfig(aws.Config{}, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("http://s3.localhost.localstack.cloud:4566")
		//o.EndpointResolver = s3.EndpointResolverFromURL("http://localhost:4566") // LocalStack endpoint
		o.UsePathStyle = true // Force path-style addressing
		o.Region = "ap-southeast-1"
		o.Credentials = credentials.NewStaticCredentialsProvider("test", "test", "")
	})
	// l := s3.ListBucketsInput{}
	// res, _ := s3Client.ListBuckets(context.TODO(), &l)

	// for _, bucket := range res.Buckets {
	// 	fmt.Printf("\t%v\n", *bucket.Name)
	// }
	return &S3Client{
		s3Client,
	}
}

func (s3Client *S3Client) S3PutObject(ctx context.Context, bucket string, key string, file io.Reader) error {
	Object := s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	}
	_, err := s3Client.PutObject(ctx, &Object)
	return err
}

func (s3Client *S3Client) S3GetObject(ctx context.Context, bucket string, key string) error {
	Object := s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}
	_, err := s3Client.GetObject(ctx, &Object)

	//fmt.Println(*res)
	return err
}

func (s3Client *S3Client) S3DeleteObject(ctx context.Context, bucket string, key string) error {
	Object := s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	_, err := s3Client.DeleteObject(ctx, &Object)

	return err
}
