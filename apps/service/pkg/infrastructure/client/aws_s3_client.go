package client

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type AwsS3Client interface {
	PutObject(ctx context.Context, bucket string, folder_name string, file []byte) (*string, error)
}

type AwsS3ClientImpl struct {
	s3Client *s3.Client
}

func NewAwsS3Client() AwsS3Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	client := s3.NewFromConfig(cfg)
	return &AwsS3ClientImpl{
		s3Client: client,
	}
}

func (c *AwsS3ClientImpl) PutObject(ctx context.Context, bucket string, folder_name string, file []byte) (*string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	fileName := fmt.Sprintf("%s/%s", folder_name, uuid.String())
	body := bytes.NewReader(file)
	input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &fileName,
		Body:   body,
		// ACL:    types.ObjectCannedACLPublicRead,
	}
	_, err = c.s3Client.PutObject(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to upload: %v", err)
	}
	region := c.s3Client.Options().Region
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName)

	return &url, nil
}
